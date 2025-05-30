package database

import (
	"log"

	"github.com/kunals12/go-todo-api/models"
)

func Migrate() {
	// Ensure the uuid-ossp extension exists
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	err := DB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Database migrated successfully")
}
