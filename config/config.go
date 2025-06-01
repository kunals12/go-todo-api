package config

import (
	"os"

	"github.com/rs/zerolog/log"
)

// Config holds application-level configuration loaded from environment variables
type Config struct {
	DBUrl        string // PostgreSQL database connection string
	Port         string // Port to run the server on
	Env          string // Application environment (e.g., development, production)
	JwtSecretKey string
}

// AppConfig is the global instance of Config used throughout the app
var AppConfig Config

// LoadEnv loads environment variables and populates AppConfig
func LoadEnv() {
	// Try to load variables from .env file (optional; defined in env.go)
	if err := loadDotEnv(); err != nil {
		log.Warn().Msg(".env file not found, using system env vars")
	}

	// Assign values to AppConfig from environment variables with fallbacks
	AppConfig = Config{
		DBUrl:        getEnv("DB_URL", ""),         // No fallback: must be provided
		Port:         getEnv("PORT", "3000"),       // Default to port 3000 if not set
		Env:          getEnv("ENV", "development"), // Default to "development" environment
		JwtSecretKey: getEnv("JWT_SECRET_KEY", ""),
	}

	// Ensure DB_URL is provided, otherwise terminate the app
	if AppConfig.DBUrl == "" {
		log.Error().Msg("DB_URL is required")
	}

	if AppConfig.JwtSecretKey == "" {
		log.Error().Msg("JWT Secret Not Found")
	}
}

// getEnv retrieves an environment variable or returns a fallback if not found
func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
