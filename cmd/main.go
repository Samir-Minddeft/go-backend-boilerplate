package main

import (
	"github.com/Samir-Minddeft/go-backend-boilerplate/config"
	"github.com/Samir-Minddeft/go-backend-boilerplate/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()

	// DB setup
	config.Connect()

	// Route setup
	router := routes.UserRoute()
	router.Run(":3000")
}
