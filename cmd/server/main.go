package main

import (
	"log"

	"github.com/Hemanth5603/resume-go-server/configs"
	"github.com/Hemanth5603/resume-go-server/internal/api/routes"
	"github.com/Hemanth5603/resume-go-server/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect to database
	_, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	app := fiber.New()

	// Register routes
	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(cfg.Port))
}
