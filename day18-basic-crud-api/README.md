# Day 18: Basic CRUD API with Gin

This directory explores how to build a scalable REST API using the **Gin** framework, following a structured **Controller-Service-Model (CSM)** architecture.

## 📝 Key Concepts Covered

- **Modular Architecture**: Organizing code into logical packages:
  - `model`: Defines the data structures (e.g., `User` struct).
  - `service`: Contains the business logic and data manipulation.
  - `controller`: Handles HTTP requests, extracts parameters, and sends responses.
  - `routes`: Centralizes route registration to keep `main.go` clean.
- **Gin Framework Advanced Usage**:
  - `ShouldBindJSON()`: Automatically mapping request bodies to Go structs with error handling.
  - `Query()` and `Param()`: Handling URL parameters and query strings.
- **HTTP Status Codes**: Consistent use of standard status codes like `201 Created`, `204 No Content`, and `404 Not Found`.

## 📂 Files

- [main.go](main.go): Server entry point and global configuration.
- [routes/](routes/): Route registration logic.
- [users/](users/): Full implementation of user-related CRUD operations including controller, service, and model.

## 🚀 How to Run

From the project root:
```bash
go run day18-basic-crud-api/main.go
```
