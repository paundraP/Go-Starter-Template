package user

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/entities"
	"Go-Starter-Template/internal/utils"
	emailservice "Go-Starter-Template/internal/utils/mailing"
	"Go-Starter-Template/internal/utils/storage"
	"Go-Starter-Template/pkg/jwt"
	"bytes"
	"context"
	"github.com/google/uuid"
	"html/template"
	"os"
	"strings"
	"time"
)

type (
	UserService interface {
		Register(ctx context.Context, req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
		Login(ctx context.Context, req domain.UserLoginRequest) (domain.UserLoginResponse, error)
		SendVerificationEmail(ctx context.Context, req domain.SendVerifyEmailRequest) error
		VerifyEmail(ctx context.Context, req domain.VerifyEmailRequest) (domain.VerifyEmailResponse, error)
		Me(ctx context.Context, userID string) (domain.DetailUserResponse, error)
		Update(ctx context.Context, req domain.UpdateUserRequest, userID string) (domain.UpdateUserResponse, error)
	}

	userService struct {
		userRepository UserRepository
		jwtService     jwt.JWTService
		S3             storage.AwsS3
	}
)

func NewUserService(userRepository UserRepository, jwtService jwt.JWTService, s3 storage.AwsS3) UserService {
	return &userService{
		userRepository: userRepository,
		jwtService:     jwtService,
		S3:             s3,
	}
}

var VerifyEmailRoute = "api/v1/users/verify"

func (s *userService) Register(ctx context.Context, req domain.UserRegisterRequest) (domain.UserRegisterResponse, error) {
	// checking user if exist
	if ok, err := s.userRepository.GetEmail(ctx, req.Email); err != nil {
		return domain.UserRegisterResponse{}, err
	} else if ok != nil {
		return domain.UserRegisterResponse{}, domain.ErrEmailAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return domain.UserRegisterResponse{}, err
	}

	user := entities.User{
		ID:       uuid.New(),
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Contact:  req.Contact,
		Role:     domain.RoleUser,
	}

	draftEmail, err := s.makeVerificationEmail(req.Email)
	if err != nil {
		return domain.UserRegisterResponse{}, err
	}

	if err := emailservice.SendMail(req.Email, draftEmail["subject"], draftEmail["body"]); err != nil {
		return domain.UserRegisterResponse{}, err
	}

	if err := s.userRepository.CreateUser(ctx, &user); err != nil {
		return domain.UserRegisterResponse{}, domain.ErrRegisterUserFailed
	}

	return domain.UserRegisterResponse{
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (s *userService) Login(ctx context.Context, req domain.UserLoginRequest) (domain.UserLoginResponse, error) {
	// check email if exist
	user, err := s.userRepository.GetEmail(ctx, req.Email)
	if err != nil {
		return domain.UserLoginResponse{}, domain.CredentialInvalid
	}
	if !user.Verified {
		return domain.UserLoginResponse{}, domain.ErrUserNotVerified
	}
	if ok := utils.CheckPassword(req.Password, user.Password); !ok {
		return domain.UserLoginResponse{}, domain.CredentialInvalid
	}

	token := s.jwtService.GenerateTokenUser(user.ID.String(), user.Role)

	return domain.UserLoginResponse{
		Token: token,
		Role:  user.Role,
	}, nil
}

func (s *userService) makeVerificationEmail(email string) (map[string]string, error) {
	expired := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	plainText := email + "_" + expired
	token, err := utils.AESEncrypt(plainText)
	if err != nil {
		return nil, err
	}
	// for this, better you use your frontend url that will fetch this link.
	// THIS IS ONLY EXAMPLE! DONT DO THIS IN PRODUCTION
	appUrl := os.Getenv("APP_URL")
	verifyLink := appUrl + "/" + VerifyEmailRoute + "?token=" + token

	readHtml, err := os.ReadFile("internal/utils/mailing/verification_mail.html")
	if err != nil {
		return nil, err
	}
	data := struct {
		Email  string
		Verify string
	}{
		Email:  email,
		Verify: verifyLink,
	}

	tmpl, err := template.New("custom").Parse(string(readHtml))
	if err != nil {
		return nil, err
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, err
	}

	draftEmail := map[string]string{
		"subject": "Verification Email",
		"body":    strMail.String(),
	}

	return draftEmail, nil
}

func (s *userService) SendVerificationEmail(ctx context.Context, req domain.SendVerifyEmailRequest) error {
	user, err := s.userRepository.GetEmail(ctx, req.Email)
	if err != nil {
		return domain.ErrEmailNotFound
	}

	draftEmail, err := s.makeVerificationEmail(user.Email)
	if err != nil {
		return err
	}
	if err := emailservice.SendMail(user.Email, draftEmail["subject"], draftEmail["body"]); err != nil {
		return err
	}
	return nil
}

func (s *userService) VerifyEmail(ctx context.Context, req domain.VerifyEmailRequest) (domain.VerifyEmailResponse, error) {
	decryptToken, err := utils.AESDecrypt(req.Token)
	if err != nil {
		return domain.VerifyEmailResponse{}, err
	}

	tokenParts := strings.Split(decryptToken, "_")
	if len(tokenParts) < 2 {
		return domain.VerifyEmailResponse{}, domain.ErrTokenInvalid
	}

	email := tokenParts[0]
	expirationDate := tokenParts[1]
	expirationTime, err := time.Parse("2006-01-02 15:04:05", expirationDate)
	if err != nil {
		return domain.VerifyEmailResponse{}, domain.ErrTokenInvalid
	}
	// email, expired, err := s.jwtService.GetUserEmailByToken(req.Token)
	// if err != nil {
	// 	return domain.VerifyEmailResponse{}, domain.ErrTokenInvalid
	// }

	now := time.Now()

	if expirationTime.Before(now) {
		return domain.VerifyEmailResponse{
			Email:      email,
			IsVerified: false,
		}, domain.ErrTokenExpired
	}

	user, err := s.userRepository.GetEmail(ctx, email)
	if err != nil {
		return domain.VerifyEmailResponse{}, domain.ErrEmailNotFound
	}

	if user.Verified {
		return domain.VerifyEmailResponse{}, domain.ErrAccountAlreadyVerified
	}

	updatedUser, err := s.userRepository.UpdateUser(ctx, entities.User{
		ID:       user.ID,
		Verified: true,
	})
	if err != nil {
		return domain.VerifyEmailResponse{}, err
	}

	return domain.VerifyEmailResponse{
		Email:      user.Email,
		IsVerified: updatedUser.Verified,
	}, nil
}

func (s *userService) Me(ctx context.Context, userID string) (domain.DetailUserResponse, error) {
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return domain.DetailUserResponse{}, domain.ErrUserNotFound
	}

	totalPoint := user.ActivePoint + user.LevelPoint
	rank, err := s.userRepository.GetRankByTotalPoint(ctx, totalPoint)
	if err != nil {
		return domain.DetailUserResponse{}, domain.ErrGetRank
	}

	return domain.DetailUserResponse{
		Name:           user.Name,
		Username:       user.Username,
		Email:          user.Email,
		Contact:        user.Contact,
		ProfilePicture: user.ProfilePicture,
		Subscription:   user.Subscribe,
		ActivePoint:    user.ActivePoint,
		LevelPoint:     user.LevelPoint,
		Rank:           rank.Name,
	}, nil
}

func (s *userService) Update(ctx context.Context, req domain.UpdateUserRequest, userID string) (domain.UpdateUserResponse, error) {
	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return domain.UpdateUserResponse{}, domain.ErrUserNotFound
	}

	if user.ProfilePicture != "" {
		updatedKey, err := s.S3.UpdateFile(s.S3.GetObjectKeyFromLink(user.ProfilePicture), req.ProfilePicture, storage.AllowImage...)
		if err != nil {
			return domain.UpdateUserResponse{}, err
		}
		user.ProfilePicture = s.S3.GetPublicLinkKey(updatedKey)
	} else if user.ProfilePicture == "" {
		objectKey, err := s.S3.UploadFile("ProfilePicture-"+user.ID.String(), req.ProfilePicture, "profile-picture", storage.AllowImage...)
		if err != nil {
			return domain.UpdateUserResponse{}, err
		}
		user.ProfilePicture = s.S3.GetPublicLinkKey(objectKey)
	}

	// validation if the user who's updating is the valid user
	id, err := uuid.Parse(userID)
	if err != nil {
		return domain.UpdateUserResponse{}, domain.ErrParseUUID
	}
	if user.ID != id {
		return domain.UpdateUserResponse{}, domain.ErrUserNotValid
	}

	user.Name = ifNotEmpty(req.Name, user.Name)
	user.Username = ifNotEmpty(req.Username, user.Username)
	user.Email = ifNotEmpty(req.Email, user.Email)
	user.Contact = ifNotEmpty(req.Contact, user.Contact)

	upd, err := s.userRepository.UpdateUser(ctx, *user)
	if err != nil {
		return domain.UpdateUserResponse{}, err
	}

	return domain.UpdateUserResponse{
		Name:           upd.Name,
		Username:       upd.Username,
		Email:          upd.Email,
		Contact:        upd.Contact,
		ProfilePicture: upd.ProfilePicture,
	}, nil
}

func ifNotEmpty(value, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
}
