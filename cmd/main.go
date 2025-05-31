package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kunals12/go-todo-api/config"
	"github.com/kunals12/go-todo-api/database"
	"github.com/kunals12/go-todo-api/routes"
)

func main() {
	config.LoadEnv()
	database.Connect()
	// database.Migrate()
	app := fiber.New()
	app.Use(recover.New())
	routes.SetupTodoRoutes(app)
	routes.UserRoutes(app)
	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	log.Fatal(app.Listen(":3000"))
}
