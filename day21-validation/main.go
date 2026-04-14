package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-basics/day21-validation/users"
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
	db.AutoMigrate(&users.User{})

	r := gin.Default()

	userService := users.NewService(db)
	userCtrl := users.NewController(userService)
	users.RegisterRoutes(r, userCtrl)

	r.Run(":5000")
}
