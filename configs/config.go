package configs

// Config holds the application configuration
import (
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

	viper.AutomaticEnv()

	// Try to read config file, but don't fail if it doesn't exist
	// Environment variables will be used instead
	_ = viper.ReadInConfig()

	err = viper.Unmarshal(&config)
	return
}
