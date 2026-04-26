package main

import (
	"fmt"
	"log"

	"go-basics/day26-environment-config/config"
	"go-basics/day26-environment-config/utils"
)

func main() {
	// Always load config first — before anything else
	config.Load()

	cfg := config.AppConfig

	fmt.Println("=== App Config ===")
	fmt.Println("Port:       ", cfg.AppPort)
	fmt.Println("Env:        ", cfg.AppEnv)
	fmt.Println("DB Host:    ", cfg.DBHost)
	fmt.Println("DB Name:    ", cfg.DBName)
	fmt.Println("JWT Expiry: ", cfg.JWTExpiry, "hours")
	fmt.Println("JWT Secret: ", "[hidden]")

	// Demo: generate and validate a token using config-driven secret
	token, err := utils.GenerateToken(1, "user@example.com", "admin")
	if err != nil {
		log.Fatal("GenerateToken failed:", err)
	}
	fmt.Println("\n=== JWT Demo ===")
	fmt.Println("Token:", token)

	claims, err := utils.ValidateToken(token)
	if err != nil {
		log.Fatal("ValidateToken failed:", err)
	}
	fmt.Printf("Validated — UserID: %d, Email: %s, Role: %s\n",
		claims.UserID, claims.Email, claims.Role)

	// ConnectDB(cfg) would go here in a real app
	// config.ConnectDB(cfg)

	fmt.Printf("\nServer starting on :%s\n", cfg.AppPort)
}
