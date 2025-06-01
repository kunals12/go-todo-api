package database

import (
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	log.Info().Msg("Connecting to DB")
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
	}

	DB = db
	log.Info().Msg("Connected to DB")
}
