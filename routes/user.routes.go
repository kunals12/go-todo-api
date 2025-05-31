package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kunals12/go-todo-api/controllers"
	"github.com/kunals12/go-todo-api/middlewares"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/users")

	user.Post("/", controllers.CreateUser) //public
	user.Get("/:id", middlewares.JWTMiddleware(), controllers.GetUserById)
	// user.Patch("/:id", controllers.PatchUser)
	// user.Delete("/:id", controllers.DeleteUser)
}
