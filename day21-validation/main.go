package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GORM model — maps to DB
type User struct {
	ID        uint           `json:"id"         gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
	Email     string         `json:"email"      gorm:"uniqueIndex"`
	Age       int            `json:"age"`
}

// Input struct — used only for request binding + validation

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,min=1,max=120"`
}

type UpdateUserInput struct {
	Name  string `json:"name"  binding:"omitempty,min=2,max=50"`
	Email string `json:"email" binding:"omitempty,email"`
	Age   int    `json:"age"   binding:"omitempty,gte=1,lte=120"`
}

var DB *gorm.DB

func connectDB() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env", err)
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("✅ GORM connected to PostgreSQL!")
}

func formatErrors(err error) map[string]string {
	errs := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errs[e.Field()] = e.Tag()
	}
	return errs
}

func createUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input validation error", "errors": formatErrors(err)})
		return
	}

	user := User{Name: input.Name, Age: input.Age, Email: input.Email}
	if result := DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func deleteUser(c *gin.Context) {
	var user User
	if result := DB.First(&user, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func main() {
	connectDB()
	DB.AutoMigrate(&User{})
	r := gin.Default()
	r.GET("/users", func(ctx *gin.Context) {
		var users []User
		DB.Find(&users)
		ctx.JSON(http.StatusOK, users)
	})

	r.POST("/users", createUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":5000")
}
