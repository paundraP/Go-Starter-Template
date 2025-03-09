package user

import (
	"Go-Starter-Template/internal/utils"
	"Go-Starter-Template/internal/utils/storage"
	"Go-Starter-Template/pkg/entities"
	"Go-Starter-Template/pkg/entities/domain"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
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
	isRegister := s.userRepository.CheckUser(ctx, req.Email)
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
