# Day 28: Testing in Go

This directory covers the fundamentals of writing unit tests in Go, focusing on testing business logic (services) using the standard testing package and popular assertion libraries.

## 📝 Key Concepts Covered

- **Standard Testing Package**: Using Go's built-in `testing` package to write test functions (`func TestXxx(t *testing.T)`).
- **Table-Driven Tests**: A standard Go pattern to run multiple test cases through the same logic cleanly, defining inputs and expected outputs in a slice of structs.
- **The `testify` Toolkit**:
  - `assert`: For clear and concise checks (e.g., `assert.Equal`, `assert.Contains`). Does not stop execution on failure.
  - `require`: Similar to `assert`, but immediately halts the test upon failure (e.g., `require.NoError`).
- **Mocking Dependencies**: Creating in-memory mocks (e.g., `mockUserRepo`) that satisfy interfaces used by the service layer. This isolates the service logic from the actual database.
- **Test Helpers**: Using `t.Helper()` inside custom utility functions to ensure that test failure line numbers point to the test code, not the helper function itself.
- **Test Coverage**: Using Go CLI flags to measure how much code is exercised by tests.

## 📂 Files

- [main.go](main.go): Provides instructions on how to run the tests.
- [domain/](domain/): Interfaces and structs defining the core entities.
- [service/user_service.go](service/user_service.go): The business logic being tested.
- [service/mock_user_repo_test.go](service/mock_user_repo_test.go): A simple in-memory implementation of the repository interface.
- [service/user_service_test.go](service/user_service_test.go): The comprehensive test suite using standard tables and `testify`.

## 🚀 How to Run

Navigate to this directory or the project root and run:

```bash
# Run all tests with verbose output
go test -v ./day28-testing/service/...

# Run tests and show coverage percentage
go test -cover ./day28-testing/service/...

# Run a specific test
go test -v -run TestUserService_Login ./day28-testing/service/...
```
