package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kunals12/go-todo-api/controllers"
)

func SetupTodoRoutes(app *fiber.App) {
	// api := app.Group("/api")
	todo := app.Group("/todos")

	todo.Get("/", controllers.GetTodos)
	todo.Post("/", controllers.CreateTodo)
	todo.Patch("/:id", controllers.PatchTodo)
	todo.Delete("/:id", controllers.DeleteTodo)
}
