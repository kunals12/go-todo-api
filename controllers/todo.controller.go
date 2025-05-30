package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// PatchTodo partially updates a todo item
func PatchTodo(c *fiber.Ctx) error {
	// 1. Parse ID from URL param
	id := c.Params("id")
	todoID, error := uuid.Parse(id)

	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}

	// 2. Find the Todo By Id
	var todo models.Todo

	if err := database.DB.First(&todo, "id = ?", todoID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	// 3. Define input struct with pointer fields
	type UpdateTodoInput struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	var input UpdateTodoInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// 4. Conditionally update the fields
	if input.Title != nil {
		todo.Title = *input.Title
	}
	if input.Completed != nil {
		todo.Completed = *input.Completed
	}

	if err := database.DB.Save(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update todo",
		})
	}

	return c.JSON(todo)
}

// Delete todo item
func DeleteTodo(c *fiber.Ctx) error {
	// Get the Id from param
	id := c.Params("id")

	todoId, error := uuid.Parse(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}

	// Attempt to delete the todo
	result := database.DB.Delete(&models.Todo{}, "id = ?", todoId)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete todo",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	// Success
	return c.SendStatus(fiber.StatusNoContent) // 204
}
