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

type ()
