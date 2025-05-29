package config

import (
	"log"
	"os"
)

type Config struct {
	DBUrl string
	Port  string
	Env   string
}

var AppConfig Config

func LoadEnv() {
	if err := loadDotEnv(); err != nil {
		log.Println("Warning: .env file not found, using system env vars")
	}

	AppConfig = Config{
		DBUrl: getEnv("DB_URL", ""),
		Port:  getEnv("PORT", "3000"),
		Env:   getEnv("ENV", "development"),
	}

	if AppConfig.DBUrl == "" {
		log.Fatal("DB_URL is required")
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
