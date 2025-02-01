package middleware

import "github.com/gofiber/fiber/v2"

type (
	Middleware interface {
		AuthMiddleware() fiber.Handler
		CORSMiddleware() fiber.Handler
		OnlyAllow(allow string) fiber.Handler
	}
	middleware struct {
	}
)

func NewMiddleware() Middleware {
	return &middleware{}
}
