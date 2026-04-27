# Day 26: Environment Configuration Management

This directory focuses on **Configuration Management**, exploring how to transition from hardcoded strings and loose environment variables to a robust, type-safe config system.

## 📝 Key Concepts Covered

- **Centralized Configuration**: Using a single `Config` struct to hold all application settings (ports, DB credentials, JWT secrets).
- **Environment Variable Fallbacks**: Implementing a `getEnv` helper that provides sensible defaults for non-sensitive values while ensuring sensitive values remain unset if not provided.
- **Boot-time Validation**: Implementing a `validate()` function that checks for missing required variables (like `JWT_SECRET` or `DB_NAME`) at startup, preventing "silent" failures later.
- **Type Safety**: Converting string environment variables to appropriate types (e.g., `int` for `JWT_EXPIRY_HOURS`) during the loading phase.
- **Clean Main Logic**: Keeping `main.go` focused on orchestration by delegating all environment parsing and error handling to the `config` package.

## 📂 Files

- [main.go](main.go): Orchestrates the application and demonstrates accessing the global config.
- [config/config.go](config/config.go): The core logic for loading, parsing, and validating environment settings.
- [utils/jwt.go](utils/jwt.go): Demonstrates how other packages can consume the centralized config (e.g., using `AppConfig.JWTSecret`).

## 🚀 How to Run

1.  Create a `.env` file with the required variables:
    ```env
    DB_HOST=localhost
    DB_USER=postgres
    DB_NAME=go_basics
    JWT_SECRET=supersecret
    ```
2.  Run the application:
    ```bash
    go run day26-environment-config/main.go
    ```
3.  Observe the startup validation. If you comment out `JWT_SECRET` in `.env`, the app will fail with a clear error message.
