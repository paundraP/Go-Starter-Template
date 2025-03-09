package domain

import (
	"errors"
	"mime/multipart"
)

var (
	MessageSuccessRegister             = "register success"
	MessageSuccessLogin                = "login success"
	MessageSuccessVerify               = "verify email success"
	MessageSuccessGetDetail            = "success get detail"
	MessageSuccessSendVerificationMail = "send verify email success"
	MessageSuccessUpdateUser           = "update user success"

	MessageFailedBodyRequest = "body request failed"
	MessageFailedRegister    = "register failed"
	MessageFailedLogin       = "login failed"
	MessageFailedGetDetail   = "failed get detail"
	MessageFailedUpdateUser  = "failed update user"

	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrEmailAlreadyExists     = errors.New("email already exists")
	ErrEmailNotFound          = errors.New("email not found")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserNotValid           = errors.New("user is not valid")
	CredentialInvalid         = errors.New("credential invalid")
	ErrUserNotVerified        = errors.New("user not verified")
	ErrRegisterUserFailed     = errors.New("register user failed")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrUploadFile             = errors.New("upload file failed")
)

type (
	UserRegisterRequest struct {
		Name           string                `json:"name" form:"name" validate:"required"`
		Password       string                `json:"password" form:"password" validate:"required"`
		Email          string                `json:"email" form:"email" validate:"required,email"`
		About          string                `json:"about" form:"about" validate:"required"`
		Address        string                `json:"address" form:"address" validate:"required"`
		CurrentTitle   string                `json:"current_title" form:"current_title"`
		ProfilePicture *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
		Headline       *multipart.FileHeader `json:"headline" form:"headline"`
	}

	UserRegisterResponse struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		About          string `json:"about"`
		Address        string `json:"address"`
		CurrentTitle   string `json:"current_title"`
		ProfilePicture string `json:"profile_picture"`
		Headline       string `json:"headline"`
		IsPremium      bool   `json:"is_premium"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	UserLoginResponse struct {
		Email string `json:"email"`
		Token string `json:"token"`
		Role  string `json:"role"`
	}

	UpdateUserRequest struct {
		Name           string                `json:"name" form:"name"`
		Email          string                `json:"email" form:"email" validate:"required,email"`
		NewEmail       string                `json:"new_email" form:"new_email" validate:"email"`
		About          string                `json:"about" form:"about"`
		Address        string                `json:"address" form:"address"`
		CurrentTitle   string                `json:"current_title" form:"current_title"`
		ProfilePicture *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
		Headline       *multipart.FileHeader `json:"headline" form:"headline"`
	}
)
