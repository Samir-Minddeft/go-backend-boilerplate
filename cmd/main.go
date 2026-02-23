package main

import (
	"os"

	"github.com/Samir-Minddeft/go-backend-boilerplate/api/router"
	"github.com/Samir-Minddeft/go-backend-boilerplate/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()

	// DB setup
	config.Connect()

	//Router setup
	r := gin.Default()
	router.Router(r)

	// Get port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)

}
