package main

import (
	"log"

	"github.com/Hemanth5603/resume-go-server/configs"
	"github.com/Hemanth5603/resume-go-server/internal/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app := fiber.New()

	// Register routes
	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(cfg.Port))
}
