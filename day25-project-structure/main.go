package main

import (
	"go-basics/day25-project-structure/config"
	"go-basics/day25-project-structure/domain"
	"go-basics/day25-project-structure/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(&domain.User{}, &domain.Post{}, &domain.PostAnalytic{})

	h := wireHandlers()

	r := gin.New()
	routes.SetupRoutes(r, h)
	r.Run(":5000")
}
