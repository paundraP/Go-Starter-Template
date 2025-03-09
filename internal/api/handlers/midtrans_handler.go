package handlers

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/internal/api/presenters"
	"Go-Starter-Template/pkg/midtrans"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	MidtransHandler interface {
		CreateTransaction(c *fiber.Ctx) error
		MidtransWebhookHandler(c *fiber.Ctx) error
	}
	midtransHandler struct {
		midtransService midtrans.MidtransService
		Validator       *validator.Validate
	}
)

func NewMidtransHandler(midtransService midtrans.MidtransService, validator *validator.Validate) MidtransHandler {
	return &midtransHandler{
		midtransService: midtransService,
		Validator:       validator,
	}
}

func (h *midtransHandler) CreateTransaction(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req domain.MidtransPaymentRequest

	if err := c.BodyParser(&req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedCreateTransaction, err)
	}
	res, err := h.midtransService.CreateTransaction(req, userID)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedCreateTransaction, err)
	}

	return presenters.SuccessResponse(c, res, fiber.StatusOK, domain.MessageSuccessCreateTransaction)
}

func (h *midtransHandler) MidtransWebhookHandler(c *fiber.Ctx) error {
	var notification domain.MidtransWebhookRequest

	if err := c.BodyParser(&notification); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.Validator.Struct(notification); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedRegister, err)
	}

	res, err := h.midtransService.MidtransWebHook(c.Context(), notification)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusOK, domain.MessageSuccessWebhook)
}
