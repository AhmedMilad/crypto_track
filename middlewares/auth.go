package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"crypto_track/database"
	"crypto_track/models"
	"crypto_track/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

	return func(c *gin.Context) {
		respondInvalidToken := func() {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Message: "Invalid token",
			})
			c.Abort()
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			respondInvalidToken()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			respondInvalidToken()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			respondInvalidToken()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			respondInvalidToken()
			return
		}

		email, ok := claims["email"].(string)
		if !ok || email == "" {
			respondInvalidToken()
			return
		}

		var user models.User
		err = database.Conn.Where("email = ?", email).First(&user).Error
		if err != nil {
			respondInvalidToken()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
