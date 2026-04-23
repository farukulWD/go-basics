package main

import (
	"go-basics/day25-project-structure/config"
	"go-basics/day25-project-structure/domain"
	"go-basics/day25-project-structure/handler"
	"go-basics/day25-project-structure/repository"
	"go-basics/day25-project-structure/routes"
	"go-basics/day25-project-structure/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(&domain.User{})

	// Wire up layers: repo -> service -> handler
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.New()
	routes.SetupRoutes(r, userHandler)
	r.Run(":5000")
}
