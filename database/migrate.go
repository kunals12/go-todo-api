package database

import "github.com/kunals12/go-todo-api/models"

func Migrate() {
	DB.AutoMigrate(&models.Todo{})
}
