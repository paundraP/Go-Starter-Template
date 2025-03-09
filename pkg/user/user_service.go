package user

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/entities"
	"Go-Starter-Template/internal/utils"
	"Go-Starter-Template/internal/utils/storage"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
		Login(ctx context.Context, req domain.UserLoginRequest) (domain.UserLoginResponse, error)
		UpdateProfile(ctx context.Context, req domain.UpdateUserRequest) error
		UpdateEducation(ctx context.Context, req domain.UpdateUserEducationRequest, userID string) error
	}

	userService struct {
		userRepository UserRepository
		awsS3          storage.AwsS3
	}
)

func NewUserService(userRepository UserRepository, awsS3 storage.AwsS3) UserService {
	return &userService{userRepository: userRepository, awsS3: awsS3}
}

var VerifyEmailRoute = "api/verify_email/user"

func (s *userService) RegisterUser(ctx context.Context, req domain.UserRegisterRequest) (domain.UserRegisterResponse, error) {
	isRegister := s.userRepository.CheckUserByEmail(ctx, req.Email)
	if isRegister {
		return domain.UserRegisterResponse{}, domain.ErrEmailAlreadyExists
	}
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return domain.UserRegisterResponse{}, err
	}

	var profilePicture, headline string
	userID := uuid.New()
	allowedMimetype := []string{"image/jpeg", "image/png"}
	if req.ProfilePicture != nil {
		filename := fmt.Sprintf("ProfilePicture-%s", userID)
		objectKey, err := s.awsS3.UploadFile(filename, req.ProfilePicture, "profile-picture", allowedMimetype...)
		if err != nil {
			return domain.UserRegisterResponse{}, domain.ErrUploadFile
		}
		profilePicture = s.awsS3.GetPublicLinkKey(objectKey)
	}

	if req.Headline != nil {
		filename := fmt.Sprintf("Headline-%s", userID)
		objectKey, err := s.awsS3.UploadFile(filename, req.Headline, "headline", allowedMimetype...)
		if err != nil {
			return domain.UserRegisterResponse{}, domain.ErrUploadFile
		}
		headline = s.awsS3.GetPublicLinkKey(objectKey)
	}

	user := entities.User{
		ID:             userID,
		Name:           req.Name,
		Password:       password,
		Email:          req.Email,
		About:          req.About,
		Address:        req.Address,
		CurrentTitle:   req.CurrentTitle,
		ProfilePicture: profilePicture,
		Headline:       headline,
		IsPremium:      false,
		Role:           domain.RoleUser,
	}

	create, err := s.userRepository.RegisterUser(ctx, user)
	if err != nil {
		return domain.UserRegisterResponse{}, domain.ErrRegisterUserFailed
	}
	return domain.UserRegisterResponse{
		Name:           create.Name,
		Email:          create.Email,
		About:          create.About,
		Address:        create.Address,
		CurrentTitle:   create.CurrentTitle,
		ProfilePicture: create.ProfilePicture,
		Headline:       create.Headline,
		IsPremium:      create.IsPremium,
	}, nil
}

func (s *userService) Login(ctx context.Context, req domain.UserLoginRequest) (domain.UserLoginResponse, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return domain.UserLoginResponse{}, domain.ErrUserNotFound
	}
	if ok := utils.CheckPassword(req.Password, user.Password); !ok {
		return domain.UserLoginResponse{}, domain.CredentialInvalid
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return domain.UserLoginResponse{}, err
	}

	return domain.UserLoginResponse{
		Email: user.Email,
		Token: token,
		Role:  user.Role,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, req domain.UpdateUserRequest) error {
	exist := s.userRepository.CheckUserByEmail(ctx, req.Email)
	if !exist {
		return domain.ErrUserNotFound
	}
	user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.NewEmail
	}
	if req.About != "" {
		user.About = req.About
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.CurrentTitle != "" {
		user.CurrentTitle = req.CurrentTitle
	}
	allowedMimetype := []string{"image/jpeg", "image/png"}
	if req.ProfilePicture != nil {

		_, err := s.awsS3.UploadFile(user.ProfilePicture, req.ProfilePicture, "profile-picture", allowedMimetype...)
		if err != nil {
			return domain.ErrUploadFile
		}

	}

	if req.Headline != nil {
		_, err := s.awsS3.UploadFile(user.Headline, req.Headline, "headline", allowedMimetype...)
		if err != nil {
			return domain.ErrUploadFile
		}
	}
	if err := s.userRepository.UpdateProfile(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateEducation(ctx context.Context, req domain.UpdateUserEducationRequest, userID string) error {
	if exist := s.userRepository.CheckUserByID(ctx, userID); !exist {
		return domain.ErrUserNotFound
	}
	userid, err := uuid.Parse(userID)
	if err != nil {
		return domain.ErrParseUUID
	}
	education := entities.UserEducation{
		ID:           uuid.New(),
		UserID:       userid,
		SchoolName:   req.SchoolName,
		Degree:       req.Degree,
		FieldOfStudy: req.FieldOfStudy,
		Description:  req.Description,
	}

	if err := s.userRepository.UpdateEducation(ctx, education); err != nil {
		return domain.ErrUpdateEducation
	}
	return nil
}
