package configs

// Config holds the application configuration
type Config struct {
	Port string
}

// LoadConfig loads configuration from environment variables or a config file
func LoadConfig() (*Config, error) {
	// For now, we'll hardcode the config
	return &Config{
		Port: ":3000",
	}, nil
}
