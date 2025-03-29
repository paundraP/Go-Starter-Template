package middleware

import (
	"Go-Starter-Template/internal/api/presenters"
	"Go-Starter-Template/pkg/entities/domain"
	jwtService "Go-Starter-Template/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (m *middleware) AuthMiddleware(jwtService jwtService.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return presenters.ErrorResponse(c, fiber.StatusUnauthorized, domain.MessageFailedProcessRequest, domain.ErrTokenNotFound)
		}
		if !strings.Contains(authHeader, "Bearer") {
			return presenters.ErrorResponse(c, fiber.StatusUnauthorized, domain.MessageFailedProcessRequest, domain.ErrTokenNotFound)
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)

		userId, userRole, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			return presenters.ErrorResponse(c, fiber.StatusUnauthorized, domain.MessageFailedProcessRequest, err)
		}
		if userId == "" {
			return presenters.ErrorResponse(c, fiber.StatusUnauthorized, domain.MessageFailedProcessRequest, err)
		}
		c.Locals("user_id", userId)
		c.Locals("role", userRole)
		c.Locals("token", authHeader)
		return c.Next()
	}
}
