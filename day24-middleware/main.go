package main

import (
	"go-basics/day24-middleware/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.MyMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.Run(":5000")
}
