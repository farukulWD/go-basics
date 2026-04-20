package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-basics/day24-middleware/config"
	"go-basics/day24-middleware/models"
)

func GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user.ToResponse())
}

type UpdateProfileInput struct {
	Name string `json:"name" binding:"omitempty,min=2,max=50"`
}

func UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, user.ToResponse())
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	var response []models.UserResponse
	for _, u := range users {
		response = append(response, u.ToResponse())
	}

	c.JSON(http.StatusOK, response)
}
