# Day 7: Structs and Methods in Go

This directory explores how to create custom data types using **Structs** and how to add functionality to them using **Methods**.

## 📝 Key Concepts Covered

### Structs
- **Definition**: Using `type Name struct { fields }` to group different data types into a single named record.
- **Initialization**: 
  - Using the `var` keyword and dot notation (`p.Name = "John"`).
  - Using struct literals (`Person{Name: "Jane", Age: 25}`).
- **Use Case**: Creating complex real-world objects like `Person` or `Product`.

### Methods
- **Receiver Functions**: Methods are functions with a special "receiver" argument, allowing you to call the function on an instance of a struct: `func (p Receiver) MethodName()`.
- **Functionality**: Implementing logic like price calculations (`FinalPrice`) or status checks (`StockStatus`) directly on the data structure.

## 📂 Files

- [man.go](man.go): Comprehensive examples of struct declaration, method implementation, and usage scenarios for `Person` and `Product` objects.

## 🚀 How to Run

From the project root:
```bash
go run day7/man.go
```
