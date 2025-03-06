package domain

import "errors"

var (
	MessageSuccessRegister             = "register success"
	MessageSuccessLogin                = "login success"
	MessageSuccessVerify               = "verify email success"
	MessageSuccessGetDetail            = "success get detail"
	MessageSuccessSendVerificationMail = "send verify email success"
	MessageSuccessUpdateUser           = "update user success"

	MessageFailedBodyRequest = "body request failed"
	MessageFailedRegister    = "register failed"
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
)

type (
	UserRegisterRequest struct {
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required,min=3"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Contact  string `json:"contact" validate:"required"`
	}

	UserRegisterResponse struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}

	SendVerifyEmailRequest struct {
		Email string `json:"email" validate:"required,email"`
	}

	VerifyEmailRequest struct {
		Token string `json:"token" form:"token" validate:"required"`
	}

	VerifyEmailResponse struct {
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	DetailUserResponse struct {
		Name         string `json:"name"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		Contact      string `json:"contact"`
		Subscription bool   `json:"subscription"`
	}

	UpdateUserRequest struct {
		Name     string `json:"name" validate:"omitempty"`
		Username string `json:"username" validate:"omitempty,min=3"`
		Email    string `json:"email" validate:"omitempty,email"`
		Contact  string `json:"contact" validate:"omitempty"`
	}

	UpdateUserResponse struct {
		Email string `json:"email"`
	}
)
