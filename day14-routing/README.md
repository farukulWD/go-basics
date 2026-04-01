# Day 14: Advanced Routing and Middleware in Go

This directory takes our web server to the next level by implementing a full CRUD (Create, Read, Update, Delete) API with structured routing and middleware.

## 📝 Key Concepts Covered

- **The `http.ServeMux` Router**: Using a multiplexer to map multiple complex paths to different handlers efficiently.
- **RESTful Resource Patterns**: Implementing standard web methods for a `BookStore` resource:
  - `GET /books`: List all books (with optional `?genre=` filtering).
  - `POST /books`: Add a new book to the store.
  - `GET /books/:id`: Retrieve a specific book by its ID.
  - `PUT /books/:id`: Update an existing book.
  - `DELETE /books/:id`: Remove a book from the store.
- **Struct-Based Handlers**: Moving beyond simple functions by attaching handlers to structs (`BookStore`). This allows sharing state (like a database or slice) without using global variables.
- **Path Segment Parsing**: Manually splitting and parsing URL paths to handle dynamic resource IDs in the standard library.
- **Middleware Pattern**: 
  - **Concept**: Functions that wrap handlers to perform tasks before or after the main logic (e.g., logging, setting headers).
  - **Chaining**: Using a helper function to apply multiple layers of middleware (Logger + JSON Header) to our routes consistently.

## 📂 Files

- [main.go](main.go): A feature-complete BookStore API demonstrating advanced routing, middleware chaining, and in-memory data management.

## 🚀 How to Run

1.  Start the server:
    ```bash
    go run day14-routing/main.go
    ```
2.  Test the routes:
    - **List all**: `curl http://localhost:8080/books`
    - **Filter by Genre**: `curl http://localhost:8080/books?genre=tech`
    - **Add Book**: 
      ```bash
      curl -X POST -H "Content-Type: application/json" -d '{"title":"Refactoring","author":"Martin Fowler","genre":"tech"}' http://localhost:8080/books
      ```
    - **Get by ID**: `curl http://localhost:8080/books/1`
    - **Update**: `curl -X PUT -d '{"title":"Go in Action"}' http://localhost:8080/books/1`
    - **Delete**: `curl -X DELETE http://localhost:8080/books/1`
