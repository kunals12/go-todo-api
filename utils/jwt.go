package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kunals12/go-todo-api/config"
	"github.com/kunals12/go-todo-api/models"
)

func GenerateJwt(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"name":    user.Name,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // expires in 3 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JwtSecretKey))
}
