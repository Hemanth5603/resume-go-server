package configs

// Config holds the application configuration
import "os"

type Config struct {
	Port        string
	DatabaseURL string
}

// LoadConfig loads configuration from environment variables or a config file
func LoadConfig() (*Config, error) {
	// For now, we'll get config from environment variables
	// In a real application, you might use a library like Viper
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	dbURL := os.Getenv("DATABASE_URL")

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}, nil
}
