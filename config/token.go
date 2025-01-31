package config

import (
	"dco_mart/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("TOKEN_KEY")))
}
