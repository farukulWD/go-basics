# Day 17: Starting with Gin Framework

This directory marks the transition from using Go's standard library for web development to using **Gin**, one of the most popular and efficient HTTP web frameworks for Go.

## 📝 Key Concepts Covered

- **The Gin Router**: Creating a default router instance with `gin.Default()`, which includes built-in middleware for logging and crash recovery.
- **Improved Routing**: Using Gin's expressive API for defining routes (e.g., `r.GET`, `r.POST`) instead of the standard `http.HandleFunc`.
- **The `gin.Context`**: Understanding the context object that encapsulates the request and response, providing helpful methods for data extraction and response generation.
- **Convenient JSON Responses**:
  - `c.JSON()`: A simplified way to send JSON responses with status codes.
  - `gin.H`: A shortcut for `map[string]interface{}` used to create JSON payloads quickly.
- **Starting the Server**: Using `r.Run()` to start the server on `:8080` (default) with cleaner console output.

## 📂 Files

- [main.go](main.go): A basic "Hello World" style API using Gin, demonstrating a `/ping` endpoint and JSON response handling.

## 🚀 How to Run

1.  Make sure you have Gin installed:
    ```bash
    go get -u github.com/github.com/gin-gonic/gin
    ```
2.  Run the application from this directory:
    ```bash
    go run day17-gin-starting/main.go
    ```
3.  Test the endpoint:
    - **GET Ping**: `curl http://localhost:8080/ping`
