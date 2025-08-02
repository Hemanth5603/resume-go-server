package routes

import (
	"github.com/Hemanth5603/resume-go-server/internal/api/handlers"
	"github.com/Hemanth5603/resume-go-server/internal/api/middleware"
	"github.com/Hemanth5603/resume-go-server/internal/di"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(app *fiber.App, container *di.Container) {

	// Create handlers
	userHandler := handlers.NewUserHandler(container.UserService)
	subscriptionHandler := handlers.NewSubscriptionHandler(container.SubscriptionService)

	// Group routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// User routes
	v1.Get("/user", userHandler.GetUser)
	v1.Post("/user", userHandler.CreateUser)

	//Subscription routes
	v1.Post("/subscribe", middleware.ClerkAuthMiddleware(), subscriptionHandler.CreateSubscription)
	v1.Get("/is_subscribed", middleware.ClerkAuthMiddleware(), subscriptionHandler.VerifySubscription)

	// Simple health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
