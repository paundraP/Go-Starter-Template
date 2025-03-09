package handlers

import (
	"Go-Starter-Template/internal/api/presenters"
	"Go-Starter-Template/pkg/entities/domain"
	"Go-Starter-Template/pkg/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler interface {
		RegisterUser(c *fiber.Ctx) error
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

func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	req := new(domain.UserRegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	req.ProfilePicture, _ = c.FormFile("profile_picture")
	req.Headline, _ = c.FormFile("headline")

	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedRegister, err)
	}

	res, err := h.UserService.RegisterUser(c.Context(), *req)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedRegister, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusCreated, domain.MessageSuccessRegister)
}
