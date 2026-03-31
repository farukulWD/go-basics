package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Go has no try/catch. Errors are just normal return values.
// The error interface is built-in:
//
//	type error interface {
//	    Error() string
//	}
//
// nil  = success
// non-nil = something went wrong

// ── 1. Custom error type ──────────────────────────────────────────────────────
// Real-world scenario: validating a user registration form

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on '%s': %s", e.Field, e.Message)
}

func validateAge(input string) (int, error) {
	age, err := strconv.Atoi(input)
	if err != nil {
		// Wrap a custom error so callers can inspect the field name
		return 0, ValidationError{Field: "age", Message: "must be a number"}
	}
	if age < 0 || age > 120 {
		return 0, ValidationError{Field: "age", Message: "must be between 0 and 120"}
	}
	return age, nil
}

// ── 2. Wrapping errors with %w ────────────────────────────────────────────────
// Real-world scenario: loading an app config file

func loadConfig(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		// %w wraps the original error so errors.Is / errors.As still work
		return "", fmt.Errorf("loadConfig: %w", err)
	}
	return string(data), nil
}

// ── 3. Sentinel errors ────────────────────────────────────────────────────────
// Real-world scenario: looking up a user in a database (simulated)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID   int
	Name string
}

var users = map[int]User{
	1: {ID: 1, Name: "Alice"},
	2: {ID: 2, Name: "Bob"},
}

func findUser(id int) (User, error) {
	u, ok := users[id]
	if !ok {
		return User{}, fmt.Errorf("findUser(%d): %w", id, ErrUserNotFound)
	}
	return u, nil
}

// ── 4. panic + recover ────────────────────────────────────────────────────────
// Real-world scenario: a JSON parser that must not crash the server on bad input

func safeParseJSON(raw string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			// Convert panic into a regular error so the caller can handle it
			err = fmt.Errorf("parseJSON panicked: %v", r)
		}
	}()

	if raw == "" {
		panic("received empty JSON string") // simulates a bug in a library
	}
	return fmt.Sprintf("parsed: %s", raw), nil
}

// ── main ──────────────────────────────────────────────────────────────────────

func main() {
	fmt.Println("=== Day 12: Error Handling ===")
	fmt.Println()

	// 1. Custom error type + errors.As
	fmt.Println("-- 1. Form validation --")
	inputs := []string{"25", "abc", "200"}
	for _, raw := range inputs {
		age, err := validateAge(raw)
		if err != nil {
			if ve, ok := errors.AsType[ValidationError](err); ok {
				// errors.As unwraps until it finds a ValidationError
				fmt.Printf("  Bad input %q → field=%s, reason=%s\n", raw, ve.Field, ve.Message)
			}
		} else {
			fmt.Printf("  Valid age: %d\n", age)
		}
	}
	fmt.Println()

	// 2. Wrapped errors + errors.Is
	fmt.Println("-- 2. Config file loading --")
	_, err := loadConfig("/etc/myapp/config.yaml")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// errors.Is unwraps the chain to find os.ErrNotExist
			fmt.Println("  Config not found — using defaults")
		} else {
			fmt.Println("  Unexpected error:", err)
		}
	}
	fmt.Println()

	// 3. Sentinel errors
	fmt.Println("-- 3. User lookup --")
	for _, id := range []int{1, 99} {
		user, err := findUser(id)
		if err != nil {
			if errors.Is(err, ErrUserNotFound) {
				fmt.Printf("  User %d: not found (show 404 page)\n", id)
			} else {
				fmt.Println("  Unexpected DB error:", err)
			}
		} else {
			fmt.Printf("  User %d: welcome, %s!\n", id, user.Name)
		}
	}
	fmt.Println()

	// 4. panic + recover
	fmt.Println("-- 4. Safe JSON parsing --")
	cases := []string{`{"key":"value"}`, ""}
	for _, raw := range cases {
		result, err := safeParseJSON(raw)
		if err != nil {
			fmt.Println("  Recovered from panic:", err)
		} else {
			fmt.Println(" ", result)
		}
	}
}
