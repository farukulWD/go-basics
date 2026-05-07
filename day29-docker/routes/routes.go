package routes

import (
	"net/http"

	"go-clean-api/handler"
	"go-clean-api/middleware"
	"go-clean-api/repository"
	"go-clean-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUserByID)
		}
	}

	return router
}
