package domain

import (
	"errors"
	"os"
)

const (
	RoleUser = "user"
	//ROLE_ADMIN  = "admin"
	//ROLE_MENTOR = "mentor"
)

var (
	MesaageUserNotAllowed       = "user not allowed"
	MessageFailedProcessRequest = "failed to process request"

	JwtSecret = os.Getenv("JWT_SECRET")

	ErrParseUUID      = errors.New("failed to parse UUID")
	ErrUserNotAllowed = errors.New("user not allowed")
	ErrTokenNotFound  = errors.New("failed to token not found")
)
