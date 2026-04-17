package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-basics/day23-auth-user-project/config"
	"go-basics/day23-auth-user-project/models"
	"go-basics/day23-auth-user-project/routes"
)

func main() {
	godotenv.Load() // load .env file

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":5000")
}
