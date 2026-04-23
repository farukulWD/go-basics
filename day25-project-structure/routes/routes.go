package routes

import (
	"go-basics/day25-project-structure/handler"
	"go-basics/day25-project-structure/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, uh *handler.UserHandler) {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", uh.Register)
		auth.POST("/login", uh.Login)
	}

	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", uh.GetProfile)
		user.PUT("/profile", uh.UpdateProfile)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.GET("/users", uh.GetAllUsers)
	}
}
