package main

import (
	"log"

	"github.com/Hemanth5603/resume-go-server/internal/api/middleware"
	"github.com/Hemanth5603/resume-go-server/internal/api/routes"
	"github.com/Hemanth5603/resume-go-server/internal/di"
	"github.com/gofiber/fiber/v2"
)

func main() {
	container, err := di.NewContainer()
	if err != nil {
		log.Fatalf("failed to create container: %v", err)
	}

	if err := middleware.InitClerkJWKS(); err != nil {
		log.Fatalf("failed to initialize clerk jwks: %v", err)
	}

	app := fiber.New()

	// Register routes
	routes.RegisterRoutes(app, container)

	// Ensure port has colon prefix for Fiber
	port := container.Config.Port
	if port[0] != ':' {
		port = ":" + port
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(app.Listen(port))
}
