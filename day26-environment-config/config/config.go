package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	AppEnv     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	JWTExpiry  int
}

var AppConfig *Config

func Load() {
	// Silently skip if .env not found (production sets vars directly)
	godotenv.Load()

	expiry, err := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))
	if err != nil {
		expiry = 24
	}

	AppConfig = &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		AppEnv:     getEnv("APP_ENV", "development"),
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		JWTSecret:  getEnv("JWT_SECRET", ""),
		JWTExpiry:  expiry,
	}

	validate(AppConfig)
}

// getEnv reads key, falls back to default if not set.
// Never provide fallbacks for sensitive values — let validate() catch them.
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func validate(cfg *Config) {
	required := map[string]string{
		"DB_HOST":    cfg.DBHost,
		"DB_USER":    cfg.DBUser,
		"DB_NAME":    cfg.DBName,
		"JWT_SECRET": cfg.JWTSecret,
	}

	for key, val := range required {
		if val == "" {
			log.Fatalf("Missing required environment variable: %s", key)
		}
	}

	fmt.Println("Config loaded. Environment:", cfg.AppEnv)
}
