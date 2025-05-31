package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kunals12/go-todo-api/controllers"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/users")

	user.Post("/", controllers.CreateUser)
	user.Get("/:id", controllers.GetUserById)
	// user.Patch("/:id", controllers.PatchUser)
	// user.Delete("/:id", controllers.DeleteUser)
}
