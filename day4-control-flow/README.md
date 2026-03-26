# Day 4: Control Flow in Go

This directory explores how to control the execution flow of a Go program using conditional statements and loops.

## 📝 Key Concepts Covered

- **Conditional Statements**: `if`, `else if`, and `else` with simple conditions.
- **Switch Statements**: A cleaner alternative to multiple `if-else` blocks for value comparison.
- **Loops**: Go only has one loop keyword, `for`, which can be used for:
  - Classic C-style loops: `for i := 0; i < 10; i++`
  - While-style loops: `for condition {}`
  - Infinite loops: `for {}`
- **Control Keywords**: Using `break` to exit loops and `continue` to skip the current iteration.

## 📂 Files

- [main.go](main.go): A collection of functions demonstrating `if-else`, `switch`, and `for` loops in various scenarios.

## 🚀 How to Run

From the project root:
```bash
make run FILE=day4-control-flow/main.go
```
