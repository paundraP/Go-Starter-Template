package middleware

import (
	"Go-Starter-Template/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type (
	Middleware interface {
		AuthMiddleware(jwtService jwt.JWTService) fiber.Handler
		CORSMiddleware() fiber.Handler
		OnlyAllow(allow string) fiber.Handler
	}
	middleware struct {
	}
)

func NewMiddleware() Middleware {
	return &middleware{}
}
