package routes

import (
	"Go-Starter-Template/internal/api/handlers"
	"Go-Starter-Template/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App             *fiber.App
	UserHandler     handlers.UserHandler
	MidtransHandler handlers.MidtransHandler
	Middleware      middleware.Middleware
}

func (c *Config) Setup() {
	c.App.Use(c.Middleware.CORSMiddleware())
	c.GuestRoute()
	c.AuthRoute()
}

func (c *Config) GuestRoute() {
	c.App.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong, its works. please"})
	})

	// user routes
	{
		c.App.Post("/api/users/register", c.UserHandler.Register)
		c.App.Post("/api/users/login", c.UserHandler.Login)
		c.App.Post("/api/users/send_verify", c.UserHandler.SendVerificationEmail)
		c.App.Get("/api/users/verify", c.UserHandler.VerifyEmail)

		c.App.Post("/webhook/midtrans", c.MidtransHandler.MidtransWebhookHandler)
	}
}

func (c *Config) AuthRoute() {
	restricted := c.App.Group("/v1/api", c.Middleware.AuthMiddleware())

	// user
	{
		restricted.Get("/users/me", c.UserHandler.Me)
		restricted.Patch("/users/update", c.UserHandler.UpdateUser)
		restricted.Post("/users/subscribe", c.MidtransHandler.CreateTransaction)

	}

	restricted.Get("/restricted", c.Middleware.OnlyAllow("admin"), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Access granted"})
	})
}
