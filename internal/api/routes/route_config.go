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
	c.User()
	c.GuestRoute()
	c.AuthRoute()
}

func (c *Config) User() {
	user := c.App.Group("/api/user")
	{
		user.Post("/register", c.UserHandler.RegisterUser)
		user.Post("/login", c.UserHandler.Login)
		user.Post("/update", c.Middleware.AuthMiddleware(), c.UserHandler.UpdateProfile)
		user.Post("/subscribe", c.Middleware.AuthMiddleware(), c.MidtransHandler.CreateTransaction)
	}
}``

func (c *Config) GuestRoute() {
	c.App.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong, its works. please"})
	})
	c.App.Post("/webhook/midtrans", c.MidtransHandler.MidtransWebhookHandler)
}

func (c *Config) AuthRoute() {
	c.App.Get("/restricted", c.Middleware.AuthMiddleware(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Access granted"})
	})
}
