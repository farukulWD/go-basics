package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/jackc/pgx/v5"
// )

// func main() {
// 	connStr := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"
// 	conn, err := pgx.Connect(context.Background(), connStr)
// 	if err != nil {
// 		log.Fatal("Unable to connect to database:", err)
// 	}

// 	defer conn.Close(context.Background())

// 	fmt.Println("✅ Connected to PostgreSQL!")
// }

var db *pgxpool.Pool

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	connStr := os.Getenv("DATABASE_URL")

	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}
	defer db.Close()

	// Ping to verify the connection actually works
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	fmt.Println("✅ Connection pool ready!")
}
