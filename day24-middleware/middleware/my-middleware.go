package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ⬆️ Runs BEFORE the handler
		fmt.Println("Before handler")

		c.Next() // ➡️ Pass control to next middleware / handler

		// ⬇️ Runs AFTER the handler
		fmt.Println("After handler")
	}
}
