package handlers

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/internal/api/presenters"
	"Go-Starter-Template/internal/utils/pagination"
	"Go-Starter-Template/pkg/rank"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	RankHandler interface {
		Leaderboard(c *fiber.Ctx) error
	}

	rankHandler struct {
		rankService rank.RankService
		Validator   *validator.Validate
	}
)

func NewRankHandler(rankService rank.RankService, validator *validator.Validate) RankHandler {
	return &rankHandler{
		rankService: rankService,
		Validator:   validator,
	}
}

func (h *rankHandler) Leaderboard(c *fiber.Ctx) error {
	req := pagination.New(c)
	res, err := h.rankService.Leaderboard(c.Context(), req)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedGetLeaderboard, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusOK, domain.MessageSuccessGetLeaderboard)
}
