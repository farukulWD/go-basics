package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	AppEnv         string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	JWTExpiryHours int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file, reading from environment variables")
	}

	expiryHours, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))

	return &Config{
		AppPort:        getEnv("APP_PORT", "8080"),
		AppEnv:         getEnv("APP_ENV", "development"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "gouser"),
		DBPassword:     getEnv("DB_PASSWORD", "gopass"),
		DBName:         getEnv("DB_NAME", "golearn"),
		JWTSecret:      getEnv("JWT_SECRET", "dev-secret"),
		JWTExpiryHours: expiryHours,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
