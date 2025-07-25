package routes

import (
	"github.com/Hemanth5603/resume-go-server/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(app *fiber.App) {
	// Create handlers
	userHandler := handlers.NewUserHandler()

	// Group routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// User routes
	v1.Get("/user", userHandler.GetUser)

	// Simple health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
