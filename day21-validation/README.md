# Day 21: Data Validation in Go

This directory explores how to implement robust **Data Validation** in a web API using Gin's built-in binding tags.

## 📝 Key Concepts Covered

- **The `binding` Tag**: Using struct tags to define validation rules that Gin automatically enforces during JSON binding.
- **Common Validation Rules**:
  - `required`: Ensures the field is present and not empty.
  - `email`: Validates that the string follows a standard email format.
  - `min`/`max` and `gte`/`lte`: Enforces numeric ranges or string length constraints.
- **Input Structs vs. Database Models**:
  - **Input Structs** (e.g., `CreateUserInput`): Tailored for what the client sends, only including validatable and writable fields.
  - **Database Models** (e.g., `User`): Represents the full persistence layer, including auto-managed fields like `ID` and timestamps.
- **`omitempty` validation**: Allowing fields to be optional during updates while still enforcing rules if they are provided.

## 📂 Files

- [main.go](main.go): Server entry point with database and route initialization.
- [users/model.go](users/model.go): Definitions for database models and specialized input validation structs.
- [users/controller.go](users/controller.go): Handles requests and returns structured validation error messages.

## 🚀 How to Run

1.  Set up your `.env` file with `DATABASE_URL`.
2.  Run the application:
    ```bash
    go run day21-validation/main.go
    ```
3.  Test validation with invalid data:
    ```bash
    # Should fail (missing fields or invalid email)
    curl -X POST -H "Content-Type: application/json" -d '{"name": "A", "email": "not-an-email"}' http://localhost:5000/users
    ```
