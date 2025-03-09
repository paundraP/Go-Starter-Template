package middleware

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/internal/api/presenters"
	"github.com/gofiber/fiber/v2"
)

func (m *middleware) OnlyAllow(allow string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		if role == allow {
			return c.Next()
		}

		return presenters.ErrorResponse(c, fiber.StatusUnauthorized, domain.MesaageUserNotAllowed, domain.ErrUserNotAllowed)
	}
}
