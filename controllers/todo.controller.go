package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kunals12/go-todo-api/database"
	"github.com/kunals12/go-todo-api/models"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Find(&todos)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todos": todos,
	})
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse Json",
		})
	}

	database.DB.Create(&todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}
