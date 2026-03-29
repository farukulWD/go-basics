# Day 11: Interfaces in Go

This directory explores **Interfaces**, a powerful way to define behavior in Go without specifying the exact implementation.

## 📝 Key Concepts Covered

- **Interface Basics**: Defining sets of method signatures using `type Name interface { }`. Any type that implements these methods automatically "satisfies" the interface.
- **Polymorphism**: Using a single interface type to represent different concrete types (e.g., passing both `Circle` and `Rectangle` to a function that requires a `Shape`).
- **The Empty Interface (`any`)**: A special interface that can hold values of any type, similar to `Object` in other languages.
- **The `Stringer` Interface**: Implementing the `String()` method to customize how your types are printed in `fmt` functions.
- **Interface with Slices**: Storing multiple different types in a single slice of an interface type.
- **Type Handling**:
  - **Type Assertion**: Checking and extracting the original concrete type from an interface variable (`val.(Type)`).
  - **Type Switch**: A specialized `switch` statement for handling multiple potential types within an interface.

## 📂 Files

- [main.go](main.go): Comprehensive examples of shape interfaces, custom stringers, and type switching logic.

## 🚀 How to Run

From the project root:
```bash
go run day11/main.go
```
