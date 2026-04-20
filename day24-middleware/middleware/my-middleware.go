package middleware

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // process request

		duration := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		fmt.Printf("%s %s %d %s\n", method, path, status, duration)
	}
}

// gin.Default() already includes gin.Logger(). Use a custom one when you want to log to an external service like Datadog or Splunk.

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("PANIC recovered: %v\n", err)
				c.JSON(500, gin.H{"error": "Internal server error"})
				c.Abort()
			}
		}()
		c.Next()
	}
}

// gin.Default() already includes gin.Recovery(). Use a custom one when you want to log panics to an external service like Sentry.

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Reuse if client sent one, otherwise generate
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set("requestID", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

var (
	limiters = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if l, exists := limiters[ip]; exists {
		return l
	}
	// 5 requests per second, burst of 10
	l := rate.NewLimiter(5, 10)
	limiters[ip] = l
	return l
}

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := getLimiter(c.ClientIP())
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		finished := make(chan struct{})
		go func() {
			c.Next()
			finished <- struct{}{}
		}()

		select {
		case <-finished:
			// completed in time
		case <-ctx.Done():
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
			c.Abort()
		}
	}
}
