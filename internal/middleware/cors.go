package middleware

import "github.com/gofiber/fiber/v2"

func (m *middleware) CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Max-Age", "86400")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, api_key, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Expose-Headers", "Content-Length")
		c.Set("Cache-Control", "no-cache")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	}
}
