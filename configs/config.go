package configs

// Config holds the application configuration
import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JWKSURL     string `mapstructure:"JWKS_URL"` // URL for JSON Web Key Set
	JWKSIssuer  string `mapstructure:"JWKS_ISSUER"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	ForwardURL  string `mapstructure:"FORWARD_URL"` // URL for model
}

// LoadConfig loads configuration from environment variables or a config file
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
