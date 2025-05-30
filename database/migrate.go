package database

import (
	"log"

	"github.com/kunals12/go-todo-api/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Database migrated successfully")
}
