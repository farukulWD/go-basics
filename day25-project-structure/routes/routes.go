package routes

import (
	"go-basics/day25-project-structure/handler"
	"go-basics/day25-project-structure/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.Handlers) {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", h.User.Register)
		auth.POST("/login", h.User.Login)
	}

	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", h.User.GetProfile)
		user.PUT("/profile", h.User.UpdateProfile)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.GET("/users", h.User.GetAllUsers)

		admin.POST("/posts", h.Post.CreatePost)
		admin.PUT("/posts/:id", h.Post.UpdatePost)
		admin.DELETE("/posts/:id", h.Post.DeletePost)
		admin.GET("/posts/:id/analytics", h.Post.GetAnalytics)
	}

	posts := api.Group("/posts")
	{
		posts.GET("", h.Post.ListPosts)
		posts.GET("/:slug", h.Post.GetBySlug)
	}

	postsAuth := api.Group("/posts")
	postsAuth.Use(middleware.AuthMiddleware())
	{
		postsAuth.POST("/:slug/like", h.Post.LikePost)
	}
}
