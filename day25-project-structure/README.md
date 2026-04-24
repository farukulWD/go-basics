# Day 25: Domain-Driven Project Structure

This directory demonstrates a **Domain-Driven Design (DDD)** inspired project structure for Go, designed to scale with complex business requirements and multiple data entities.

## 📝 Key Concepts Covered

- **The `domain/` Package**: Centralizing all core logic for a specific entity (User, Post, Analytics) into a single directory containing:
  - **Models**: Database schemas and JSON definitions.
  - **Repositories**: Database interaction logic.
  - **Services**: Pure business logic.
  - **Handlers**: HTTP request/response logic.
- **Manual Dependency Injection**: Using a `wireHandlers` function to manually instantiate and connect models, repositories, and services without global variables.
- **Multi-Model AutoMigration**: Automatically managing schema updates for complex relationships (`User`, `Post`, `PostAnalytic`).
- **Scalable Routing**: Passing fully "wired" handlers to the route registration logic for cleaner separation of concerns.

## 📂 Files

- [main.go](main.go): The entry point that wires the entire application together.
- [domain/](domain/): Organized by feature/entity (e.g., `user.go`, `post.go`).
- [config/](config/): Shared configuration logic for databases.

## 🚀 How to Run

1.  Configure your `.env` file.
2.  Run the application:
    ```bash
    go run day25-project-structure/main.go
    ```
