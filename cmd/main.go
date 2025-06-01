package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kunals12/go-todo-api/config"
	"github.com/kunals12/go-todo-api/database"
	"github.com/kunals12/go-todo-api/routes"
)

func main() {
	config.LoadEnv()
	database.Connect()
	database.Migrate()

	app := fiber.New()
	routes.UserRoutes(app)
	routes.SetupTodoRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")
}
