package main

import (
	"fmt"
	"log"

	"go-clean-api/config"
	"go-clean-api/domain"
	"go-clean-api/routes"
)

func main() {
	cfg := config.Load()

	config.ConnectDB(cfg)
	config.DB.AutoMigrate(&domain.User{})

	router := routes.Setup(config.DB)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server starting on %s (env: %s)", addr, cfg.AppEnv)
	log.Fatal(router.Run(addr))
}
