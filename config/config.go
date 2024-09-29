package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

// Config holds all configuration values
type Config struct {
	DBConnString        string
	DefaultEmail        string
	DefaultPasswordHash string
	YubicoClientID      string
	YubicoSecretKey     string
}

// LoadConfig loads environment variables into the config struct
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	return Config{
		DBConnString:        os.Getenv("DB_CONN_STRING"),
		DefaultEmail:        os.Getenv("DEFAULT_EMAIL"),
		DefaultPasswordHash: os.Getenv("DEFAULT_PASSWORD_HASH"),
		YubicoClientID:      os.Getenv("YUBICO_CLIENT_ID"),
		YubicoSecretKey:     os.Getenv("YUBICO_SECRET_KEY"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		if defaultValue == "" {
			log.Fatal().Str("variable", key).Msg("Environment variable not set")
		}
		return defaultValue
	}
	return value
}
