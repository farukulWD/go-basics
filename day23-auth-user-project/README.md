# Day 23: Complete Auth User Project

This directory contains a professionally structured **User Authentication and Management API**. It brings together everything learned about Gin, GORM, JWT, and middleware into a scalable project architecture.

## 📝 Key Concepts Covered

- **Clean Architecture**: Organizing the project into distinct packages for better concerns separation:
  - `config`: Database connection and environment management.
  - `models`: Database schemas and GORM configurations.
  - `routes`: Centralized route definitions and grouping.
  - `handlers`: Business logic for handling requests (Login, Register, Profile).
  - `middleware`: Security layers like JWT authentication.
  - `utils`: Shared helper functions (JWT token generation).
- **Environment Management**: Robust handling of secrets and configuration using `.env` files.
- **Stateless Authentication**: Full implementation of JWT-based login and protected profile access.
- **Project Structure Best Practices**: Moving from single-file scripts to a package-based layout suitable for larger applications.
- **API Documentation**: Includes a [Postman Collection](day23-auth.postman_collection.json) to quickly test all endpoints (Register, Login, Me).

## 📂 Project Layout

- `main.go`: The clean entry point that initializes the DB and starts the router.
- `config/db.go`: Handles PostgreSQL connection pool.
- `routes/auth_routes.go`: Defines public and private route groups.
- `handlers/user_controller.go`: Contains the CRUD and Auth logic.

## 🚀 How to Run

1.  Navigate to the project directory:
    ```bash
    cd day23-auth-user-project
    ```
2.  Set up your `.env` file with `DATABASE_URL` and `JWT_SECRET`.
3.  Run the application:
    ```bash
    go run .
    ```
4.  Import [day23-auth.postman_collection.json](day23-auth.postman_collection.json) into Postman to start testing.
