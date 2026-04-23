package main

import (
	"go-basics/day25-project-structure/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()
}
