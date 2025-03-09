package handlers

import (
	"Go-Starter-Template/pkg/user"
	"github.com/go-playground/validator/v10"
)

type (
	UserHandler interface {
	}
	userHandler struct {
		UserService user.UserService
		Validator   *validator.Validate
	}
)

func NewUserHandler(userService user.UserService, validator *validator.Validate) UserHandler {
	return &userHandler{
		UserService: userService,
		Validator:   validator,
	}
}
