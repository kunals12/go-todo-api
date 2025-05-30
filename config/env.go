package config

import "github.com/joho/godotenv"

func loadDotEnv() error {
	return godotenv.Load(".env")
}
