# Day 19: PostgreSQL Connection with `pgx`

This directory explores how to connect a Go application to a **PostgreSQL** database using the modern `pgx` driver and connection pooling.

## 📝 Key Concepts Covered

- **The `pgx` Driver**: Using `github.com/jackc/pgx/v5`, a high-performance PostgreSQL driver for Go.
- **Connection Pooling**: Implementing `pgxpool` to efficiently manage multiple concurrent database connections.
- **Environment Variables**: Using `github.com/joho/godotenv` to safely load sensitive database credentials (e.g., `DATABASE_URL`) from a `.env` file.
- **Connection Validation**: Using `db.Ping()` to ensure the connection is active and ready before the application starts processing logic.
- **Context Handling**: Using `context.Background()` to manage request lifecycles and timeouts during database operations.

## 📂 Files

- [main.go](main.go): A focused example for establishing and verifying a PostgreSQL connection pool.

## 🚀 How to Run

1.  Ensure you have a PostgreSQL instance running.
2.  Create a `.env` file in the project root (or inside this directory) with your `DATABASE_URL`:
    ```env
    DATABASE_URL=postgres://username:password@localhost:5432/database_name
    ```
3.  Run the application:
    ```bash
    go run day19-postgresql-connection/main.go
    ```
