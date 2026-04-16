package main

import (
	"fmt"
	"go-basics/day22-jwt-auth/handlers"
	"go-basics/day22-jwt-auth/middleware"
	"go-basics/day22-jwt-auth/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env", err)
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("✅ GORM connected to PostgreSQL!")
	return db
}

func main() {
	db := connectDB()
	db.AutoMigrate(&models.User{})

	handlers.DB = db

	r := gin.Default()

	// Public routes
	public := r.Group("/api")
	{
		public.POST("/register", handlers.Register)
		public.POST("/login", handlers.Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			email, _ := c.Get("email")
			c.JSON(200, gin.H{"user_id": userID, "email": email})
		})
	}

	r.Run(":8080")
}
