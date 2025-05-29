package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kunals12/go-todo-api/config"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	fmt.Print(todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// todo := Todo{
	// 	ID:    1,
	// 	Title: "Do this",
	// 	Done:  true,
	// }

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "201",
		"message": "Created Todo",
		"data":    todo,
	})
}

func main() {
	config.LoadEnv()
	app := fiber.New()

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		panic("This panic is caught by fiber")
	})
	app.Post("/", createTodo)

	log.Fatal(app.Listen(":3000"))
}
