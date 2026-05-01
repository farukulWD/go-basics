# Day 27: Structured Logging with Zerolog

This directory explores how to implement high-performance, **Structured Logging** using the `github.com/rs/zerolog` package. Proper logging is essential for observability, debugging, and monitoring in production environments.

## 📝 Key Concepts Covered

- **Structured vs Unstructured Logging**: Moving from `fmt.Println` to JSON-structured logs that are machine-readable and easily searchable in log aggregators (like ELK, Datadog, Loki).
- **Environment-Specific Output**:
  - **Development**: Pretty-printed, human-readable console output with colors.
  - **Production**: High-performance JSON output, optimized for speed and parsability.
- **Log Levels**: Utilizing semantic log levels (`Trace`, `Debug`, `Info`, `Warn`, `Error`, `Fatal`) to categorize log importance.
- **Contextual Logging (Structured Fields)**: Adding specific key-value pairs (like `method`, `path`, `status`) to log entries instead of concatenating strings.
- **Child Loggers**: Creating scoped loggers that automatically inject common fields (e.g., `component: database` or `request_id`) into every subsequent log call.
- **Middleware Integration**: Building a custom Gin middleware (`middleware.RequestLogger`) that captures HTTP request details (latency, status, IP) and logs them cleanly with appropriate severity levels.
- **File Logging**: Writing logs to both the console and a physical file simultaneously using `zerolog.MultiLevelWriter`.

## 📂 Files

- [main.go](main.go): Demonstrates log levels, structured fields, child loggers, and file output.
- [config/logger.go](config/logger.go): Contains the setup logic to format logs based on the current environment.
- [middleware/logger.go](middleware/logger.go): A Gin middleware that logs HTTP request lifecycles.

## 🚀 How to Run

1.  Install zerolog:
    ```bash
    go get github.com/rs/zerolog/log
    ```
2.  Run the application:
    ```bash
    go run day27-logging/main.go
    ```
3.  Test the middleware endpoints:
    ```bash
    curl http://localhost:8080/health
    curl http://localhost:8080/bad-request
    curl http://localhost:8080/server-error
    ```
4.  Check the `app.log` file generated in the directory to see file logging in action.
