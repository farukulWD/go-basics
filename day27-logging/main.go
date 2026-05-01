package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"go-basics/day27-logging/config"
	"go-basics/day27-logging/middleware"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Logger must be set up first so all startup messages use correct format
	config.SetupLogger(env)

	demonstrateLogLevels()
	demonstrateStructuredFields()
	demonstrateChildLoggers()
	demonstrateFileLogging()

	startServer()
}

func demonstrateLogLevels() {
	log.Trace().Msg("Trace: most verbose - internal flow tracing")
	log.Debug().Msg("Debug: not shown in production")
	log.Info().Msg("Info: normal operational message")
	log.Warn().Msg("Warn: unexpected but not fatal")
	log.Error().Msg("Error: something went wrong")
	// log.Fatal() and log.Panic() omitted - they exit/panic the process
}

func demonstrateStructuredFields() {
	log.Info().
		Str("method", "POST").
		Str("path", "/api/auth/login").
		Int("status", 200).
		Dur("latency", 3*time.Millisecond).
		Msg("Request handled")

	// Simulate attaching an error
	err := os.ErrNotExist
	log.Error().
		Err(err).
		Str("file", "config.yaml").
		Msg("Config file missing")
}

func demonstrateChildLoggers() {
	// Child logger with shared "component" field - no need to repeat on every call
	dbLogger := log.With().Str("component", "database").Logger()
	dbLogger.Info().Str("query", "SELECT * FROM users").Msg("Query executed")
	dbLogger.Warn().Int("rows", 50000).Msg("Large result set")

	// Request-scoped logger - bakes request_id into every entry
	requestLogger := log.With().Str("request_id", "abc-123").Logger()
	requestLogger.Info().Msg("Processing request")
	requestLogger.Warn().Msg("Slow query detected")
}

func demonstrateFileLogging() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error().Err(err).Msg("Could not open log file")
		return
	}
	defer logFile.Close()

	// Write to both console and file simultaneously
	multi := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
		logFile,
	)
	fileLogger := zerolog.New(multi).With().Timestamp().Logger()
	fileLogger.Info().Str("component", "file-logger").Msg("This entry goes to both console and app.log")

	log.Info().Msg("File logging demo done - check app.log")
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.RequestLogger())

	// 2xx -> Info log
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 4xx -> Warn log
	r.GET("/bad-request", func(c *gin.Context) {
		log.Warn().Str("handler", "bad-request").Msg("Invalid input received")
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad input"})
	})

	// 5xx -> Error log
	r.GET("/server-error", func(c *gin.Context) {
		log.Error().Str("handler", "server-error").Msg("Simulated internal error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	})

	log.Info().Str("port", "8080").Msg("Server started - visit /health, /bad-request, /server-error")
	r.Run(":8080")
}
