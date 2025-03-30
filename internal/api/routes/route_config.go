package routes

import (
	"Go-Starter-Template/internal/api/handlers"
	"Go-Starter-Template/internal/middleware"
	"Go-Starter-Template/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App             *fiber.App
	UserHandler     handlers.UserHandler
	MidtransHandler handlers.MidtransHandler
	Middleware      middleware.Middleware
	JWTService      jwt.JWTService
}

func (c *Config) Setup() {
	c.App.Use(c.Middleware.CORSMiddleware())
	c.User()
	c.GuestRoute()
	c.AuthRoute()
}

func (c *Config) User() {
	user := c.App.Group("/api-spec/v1/users")
	// user routes
	{
		user.Post("/register", c.UserHandler.Register)
		user.Post("/login", c.UserHandler.Login)
		user.Post("/send_verify", c.UserHandler.SendVerificationEmail)
		user.Get("/verify", c.UserHandler.VerifyEmail)
		user.Get("/me", c.Middleware.AuthMiddleware(c.JWTService), c.UserHandler.Me)
		user.Patch("/update", c.Middleware.AuthMiddleware(c.JWTService), c.UserHandler.UpdateUser)
		user.Post("/subscribe", c.Middleware.AuthMiddleware(c.JWTService), c.MidtransHandler.CreateTransaction)
	}
}

func (c *Config) GuestRoute() {
	c.App.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong, its works. test"})
	})
	c.App.Post("/webhook/midtrans", c.MidtransHandler.MidtransWebhookHandler)
}

func (c *Config) AuthRoute() {
	c.App.Get("/restricted", c.Middleware.AuthMiddleware(c.JWTService), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Access granted"})
	})
	c.App.Get("/me", c.Middleware.AuthMiddleware(c.JWTService), func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		role := c.Locals("role")
		return c.JSON(fiber.Map{
			"message": "Welcome to your dashboard",
			"user_id": userID,
			"role":    role,
		})
	})
}
