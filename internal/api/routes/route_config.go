package routes

import (
	"Go-Starter-Template/internal/api/handlers"
	"Go-Starter-Template/internal/middleware"
	jwtService "Go-Starter-Template/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App             *fiber.App
	UserHandler     handlers.UserHandler
	MidtransHandler handlers.MidtransHandler
	Middleware      middleware.Middleware
	JwtService      jwtService.JWTService
}

func (c *Config) Setup() {
	c.App.Use(c.Middleware.CORSMiddleware())
	c.User()
	c.GuestRoute()
	c.AuthRoute()
}

func (c *Config) User() {
	user := c.App.Group("/api/user")
	{
		user.Post("/register", c.UserHandler.RegisterUser)
		user.Post("/login", c.UserHandler.Login)
		user.Post("/update-profile", c.Middleware.AuthMiddleware(c.JwtService), c.UserHandler.UpdateProfile)
		user.Post("/update-education", c.Middleware.AuthMiddleware(c.JwtService), c.UserHandler.UpdateEducation)
		user.Post("/subscribe", c.Middleware.AuthMiddleware(c.JwtService), c.MidtransHandler.CreateTransaction)
	}
}

func (c *Config) GuestRoute() {
	c.App.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong, its works. please"})
	})
	c.App.Post("/webhook/midtrans", c.MidtransHandler.MidtransWebhookHandler)
}

func (c *Config) AuthRoute() {
	c.App.Get("/restricted", c.Middleware.AuthMiddleware(c.JwtService), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Access granted"})
	})
	c.App.Get("/me", c.Middleware.AuthMiddleware(c.JwtService), func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		role := c.Locals("role")
		return c.JSON(fiber.Map{
			"message": "Welcome to your dashboard",
			"user_id": userID,
			"role":    role,
		})
	})
}
