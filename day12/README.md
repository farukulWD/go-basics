# Day 12: Error Handling in Go

This directory explores Go's unique approach to **Error Handling**, focusing on explicit returns instead of exceptions.

## 📝 Key Concepts Covered

- **The `error` Interface**: Understanding the built-in `error` interface and how Go uses multiple return values to signal success or failure.
- **Custom Error Types**: Creating structs that implement the `Error() string` method to provide rich, field-specific error information.
- **Sentinel Errors**: Defining constant-like error values using `errors.New()` for common failure scenarios (e.g., `ErrUserNotFound`).
- **Error Wrapping & Unwrapping**:
  - Using `fmt.Errorf` with the `%w` verb to wrap errors while preserving the original context.
  - Using `errors.Is()` to check if an error chain contains a specific sentinel error.
  - Using `errors.As()` to extract a specific custom error type from an error chain.
- **Deferred Functions, Panic, and Recover**:
  - **Panic**: When to use `panic()` for unrecoverable runtime errors.
  - **Recover**: Using `recover()` inside a `defer` block to catch a panic and convert it into a regular error.

## 📂 Files

- [main.go](main.go): Practical examples of form validation, file loading, database lookups, and safe JSON parsing using robust error handling patterns.

## 🚀 How to Run

From the project root:
```bash
go run day12/main.go
```
