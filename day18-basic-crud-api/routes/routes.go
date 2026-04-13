package routes

import (
	"go-basics/day18-basic-crud-api/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	users.RegisterUserRoutes(router)
}
