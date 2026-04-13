package main

import (
	"go-basics/day18-basic-crud-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":5000")
}
