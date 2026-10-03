package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(email string) (string, error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
