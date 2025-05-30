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
	database.Migrate()
	app := fiber.New()
	app.Use(recover.New())
	routes.SetupTodoRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
