# Day 10: Pointers in Go

This directory explores **Pointers**, an essential concept in Go for managing memory addresses and modifying variables directly from within functions.

## 📝 Key Concepts Covered

- **Memory Addresses**: Understanding that every variable is stored at a specific memory address.
- **Pointer Operators**:
  - `&` (Address-of): Used to get the memory address of a variable.
  - `*` (Dereference): Used to access or modify the value stored at a specific memory address.
- **Pass by Value vs. Pass by Reference**:
  - **Pass by Value**: Passing a copy of a variable to a function (the original remains unchanged).
  - **Pass by Reference (via Pointers)**: Passing the memory address of a variable, allowing the function to modify the original variable.
- **Pointers and Structs**: Using pointers to efficiently pass large data structures like structs and modify their fields.
- **The `new()` Keyword**: A built-in function to allocate memory for a variable and return a pointer to its address.
- **Nil Pointers**: Handling "empty" pointers to avoid runtime errors during dereferencing.

## 📂 Files

- [main.go](main.go): Detailed examples showing pointer basics, function interactions, and struct modifications.

## 🚀 How to Run

From the project root:
```bash
go run day10/main.go
```
