# Day 16: Mini User Management API

This directory contains a small, feature-complete User Management API. It demonstrates how to organize a Go project into multiple files (models, handlers, and main) to build a clean and maintainable REST API.

## 📝 Key Concepts Covered

- **Project Organization**: Splitting code into logical files:
  - `model.go`: For data structures (Structs).
  - `handler.go`: For HTTP request handling logic.
  - `main.go`: For server initialization and routing.
- **RESTful Endpoints**:
  - `GET /users`: List all users.
  - `POST /users`: Create a new user with an auto-incrementing ID.
  - `GET /users/:id`: Fetch a single user by their unique ID.
  - `DELETE /users/:id`: Remove a user from the in-memory database.
- **Advanced Path Parsing**: Using `strings.TrimPrefix` to handle dynamic URL segments in the standard library's `http.ServeMux`.
- **In-Memory State**: Managing a slice of structs globally across multiple handler functions.

## 📂 Files

- [main.go](main.go): Server setup, global state management, and route registration.
- [model.go](model.go): Defines the `User` struct with JSON tags.
- [handler.go](handler.go): Contains the implementation for `usersHandler` and `userByIDHandler`.

## 🚀 How to Run

1.  Run the API from this directory:
    ```bash
    go run .
    ```
2.  Test the endpoints:
    - **Add User**:
      ```bash
      curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@example.com"}' http://localhost:8080/users
      ```
    - **List all**: `curl http://localhost:8080/users`
    - **Get by ID**: `curl http://localhost:8080/users/1`
    - **Delete**: `curl -X DELETE http://localhost:8080/users/1`
