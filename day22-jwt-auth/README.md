# Day 22: JWT Authentication in Go

This directory covers implementing a secure **User Authentication** system using JSON Web Tokens (JWT) and the Gin framework.

## 📝 Key Concepts Covered

- **JWT Foundations**:
  - **Generation**: Creating signed tokens upon successful login to represent user identity.
  - **Verification**: Validating incoming tokens from the `Authorization` header to secure protected routes.
- **Secure Password Storage**: Using `golang.org/x/crypto/bcrypt` to hash passwords before storing them in the database, ensuring user security even if the database is compromised.
- **Authentication Middleware**:
  - Implementing custom middleware to intercept requests to protected routes.
  - Extracting and validating the `Bearer` token.
  - Injecting user data (like `userID`) into the request context (`gin.Context`) for use in downstream handlers.
- **Route Grouping**: Organizing the API into `Public` (Register/Login) and `Protected` (User Profile) groups for better structure and easier middleware application.
- **In-Memory and Persistent State**: Syncing user data with a PostgreSQL database via GORM while managing stateless session tokens.

## 📂 Files

- [main.go](main.go): Orchestrates the server, database connection, and route grouping.
- [handlers/auth.go](handlers/auth.go): Implements logic for user registration and login.
- [middleware/auth.go](middleware/auth.go): Contains the JWT validation middleware.
- [utils/jwt.go](utils/jwt.go): Helper functions for signing and parsing JWT tokens.
- [models/user.go](models/user.go): Data model for users with secure password handling.

## 🚀 How to Run

1.  Ensure you have a `.env` file with `DATABASE_URL` and `JWT_SECRET`.
2.  Install dependencies:
    ```bash
    go get github.com/golang-jwt/jwt/v5
    go get golang.org/x/crypto/bcrypt
    ```
3.  Run the application:
    ```bash
    go run day22-jwt-auth/main.go
    ```
4.  Test Authentication:
    - **Register**: `POST /api/register` with name, email, password.
    - **Login**: `POST /api/login` to receive a token.
    - **Access Protected**: `GET /api/me` with `Authorization: Bearer <token>` header.
