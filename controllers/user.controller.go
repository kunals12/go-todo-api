package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kunals12/go-todo-api/database"
	"github.com/kunals12/go-todo-api/models"
	"github.com/kunals12/go-todo-api/utils"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse Json",
		})
	}

	existingUser, err := getUserByName(user.Name)

	if err == nil {
		//User exists — generate token
		token, err := utils.GenerateJwt(*existingUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate token",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	}

	// If not found, create new user
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	token, err := utils.GenerateJwt(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func getUserByName(name string) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "name = ?", name).Error; err != nil {
		return nil, err
	}

	return &user, nil
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
