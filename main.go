package main

import (
	"crypto_track/routes/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"crypto_track/database"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.Init()

	r := gin.Default()

	api.Handle(r)

	r.Run(os.Getenv("DOMAIN"))
}
