package main

import (
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
	r.Run(":3000")

}
