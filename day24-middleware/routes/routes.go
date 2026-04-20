package routes

import (
	"time"

	"go-basics/day24-middleware/handlers"
	"go-basics/day24-middleware/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.Timeout(10 * time.Second))

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	user := api.Group("/user")
	user.Use(middleware.Auth())
	{
		user.GET("/profile", handlers.GetProfile)
		user.PUT("/profile", handlers.UpdateProfile)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.Auth(), middleware.RequireRole("admin"))
	{
		admin.GET("/users", handlers.GetAllUsers)
	}
}
