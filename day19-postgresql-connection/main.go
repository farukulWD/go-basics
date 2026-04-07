package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	connStr := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("✅ Connected to PostgreSQL!")
}
