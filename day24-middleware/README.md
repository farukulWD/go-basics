# Day 24: Middleware in Depth

This directory explores advanced **Middleware Patterns** in Go using the Gin framework, focusing on how to build production-ready applications with security and observability.

## 📝 Key Concepts Covered

- **`gin.New()` vs `gin.Default()`**: Understanding how to start with a blank slate router and manually add only the middleware you need.
- **Global Middleware**: Using `r.Use()` to apply logic to every request in the application.
- **Custom Security Middleware**:
  - **Recovery**: Catching panics to prevent server crashes and return 500 errors.
  - **CORS**: Implementing Cross-Origin Resource Sharing for frontend compatibility.
  - **Rate Limiting**: Protecting the API from brute force and denial-of-service attacks.
- **Observability Middleware**:
  - **Logger**: Recording method, path, and duration for every request.
  - **RequestID**: Injecting a unique UUID into every request context for better distributed tracing.

## 📂 Files

- [main.go](main.go): Demonstrates the registration of multiple global middlewares in a specific order.
- [middleware/](middleware/): Contains custom implementations for Rate Limiting, RequestID, and Logger.

## 🚀 How to Run

1.  Set up your `.env` file.
2.  Install the CORS package:
    ```bash
    go get github.com/gin-contrib/cors
    ```
3.  Run the server:
    ```bash
    go run day24-middleware/main.go
    ```
