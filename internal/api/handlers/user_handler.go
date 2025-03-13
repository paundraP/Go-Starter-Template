package handlers

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/internal/api/presenters"
	"Go-Starter-Template/pkg/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler interface {
		RegisterUser(c *fiber.Ctx) error
		Login(c *fiber.Ctx) error
		UpdateProfile(c *fiber.Ctx) error
		UpdateEducation(c *fiber.Ctx) error
		PostExperience(c *fiber.Ctx) error
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
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	res, err := h.UserService.RegisterUser(c.Context(), *req)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedRegister, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusCreated, domain.MessageSuccessRegister)
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	req := new(domain.UserLoginRequest)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	res, err := h.UserService.Login(c.Context(), *req)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedLogin, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusOK, domain.MessageSuccessLogin)
}

func (h *userHandler) UpdateProfile(c *fiber.Ctx) error {
	req := new(domain.UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	req.ProfilePicture, _ = c.FormFile("profile_picture")
	req.Headline, _ = c.FormFile("headline")
	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.UserService.UpdateProfile(c.Context(), *req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedUpdateUser, err)
	}

	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessUpdateUser)
}

func (h *userHandler) UpdateEducation(c *fiber.Ctx) error {
	req := new(domain.UpdateUserEducationRequest)
	userid := c.Locals("user_id").(string)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.UserService.UpdateEducation(c.Context(), *req, userid); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedAddEducation, err)
	}
	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessAddEducation)
}

func (h *userHandler) PostExperience(c *fiber.Ctx) error {
	req := new(domain.PostUserJobRequest)

	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	userid := c.Locals("user_id").(string)

	if err := h.UserService.PostJob(c.Context(), *req, userid); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedAddEducation, err)
	}

	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessAddEducation)
}
