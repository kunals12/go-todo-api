package services

import (
	"github.com/kunals12/go-todo-api/database"
	"github.com/kunals12/go-todo-api/models"
	"github.com/kunals12/go-todo-api/utils"
	"github.com/rs/zerolog/log"
)

func CreateUser(user *models.User) (string, error) {
	// Check if user already exists
	existingUser, err := GetUserByName(user.Name)

	if err != nil {
		return "", err
	}

	if existingUser != nil {
		log.Info().Msg("Existing user found")
		return utils.GenerateJwt(*existingUser)
	}

	// If not found, create new user
	if err := database.DB.Create(&user).Error; err != nil {
		return "", nil
	}

	return utils.GenerateJwt(*user)
}

func GetUserByName(name string) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "name = ?", name).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
