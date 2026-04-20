package main

import (
	"go-basics/day24-middleware/config"
	"go-basics/day24-middleware/middleware"
	"go-basics/day24-middleware/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()
	// Global — applies to ALL routes
	r := gin.New()
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.RequestID())
	r.Use(middleware.RateLimit())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	routes.SetupRoutes(r)

	r.Run(":5000")
}
