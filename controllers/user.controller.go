package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kunals12/go-todo-api/database"
	"github.com/kunals12/go-todo-api/models"
	"github.com/kunals12/go-todo-api/services"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse Json",
		})
	}

	token, err := services.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User creation failed: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	userId, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Incorrect User Id",
		})
	}
	var user models.User
	if err := database.DB.Preload("Todos").First(&user, "id = ?", userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

// func PatchUser(c *fiber.Ctx) error {

// }

// func DeleteUser(c *fiber.Ctx) error {

// }
