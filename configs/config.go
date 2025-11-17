package configs

// Config holds the application configuration
import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JWKSURL     string `mapstructure:"JWKS_URL"` // URL for JSON Web Key Set
}

// LoadConfig loads configuration from environment variables or a config file
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Bind environment variables explicitly
	viper.BindEnv("PORT")
	viper.BindEnv("DATABASE_URL")
	viper.BindEnv("JWKS_URL")

	viper.AutomaticEnv()

	// Try to read config file, but don't fail if it doesn't exist
	// Environment variables will be used instead
	_ = viper.ReadInConfig()

	err = viper.Unmarshal(&config)

	// Debug logging
	log.Printf("Config loaded - Port: %s, DatabaseURL length: %d, JWKSURL: %s",
		config.Port,
		len(config.DatabaseURL),
		config.JWKSURL)

	if len(config.DatabaseURL) > 0 {
		log.Printf("DatabaseURL starts with: %s", config.DatabaseURL[:min(50, len(config.DatabaseURL))])
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
