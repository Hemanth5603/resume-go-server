package middleware

import "github.com/gofiber/fiber/v2"

// AuthRequired is a middleware to protect routes that require authentication.
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// In a real application, you would validate a JWT or session token here.
		// For now, we'll just proceed.
		return c.Next()
	}
}
