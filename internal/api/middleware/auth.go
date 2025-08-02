package middleware

import (
	"strings"

	"github.com/Hemanth5603/resume-go-server/configs"
	"github.com/MicahParks/keyfunc"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var JWKS *keyfunc.JWKS

func InitClerkJWKS() error {
	config, err := configs.LoadConfig()
	if err != nil {
		return err
	}
	JWKS, err = keyfunc.Get(config.JWKSURL, keyfunc.Options{})
	return err
}

func ClerkAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization format"})
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, JWKS.Keyfunc)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
		}

		// Optional: validate issuer
		config, err := configs.LoadConfig()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to load config"})
		}
		if claims["iss"] != config.JWKSIssuer {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token issuer"})
		}

		// Store claims in context for access in handlers
		c.Locals("userClaims", claims)

		return c.Next()
	}
}
