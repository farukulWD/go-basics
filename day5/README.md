# Day 5: Arrays and Slices in Go

This directory covers two fundamental data structures in Go for storing collections: Arrays (fixed size) and Slices (dynamic size).

## 📝 Key Concepts Covered

### Arrays
- **Fixed Size**: Arrays have a defined length that cannot be changed once declared.
- **Declaration**: `[length]type{values}` or `[...]type{values}` to infer length.
- **Access & Modify**: Accessing elements via index `arr[0]` and modifying them.
- **Properties**: Understanding `len()` and `cap()` for arrays (both equal the array size).
- **Initialization**: Default "zero values" for uninitialized arrays.

### Slices
- **Dynamic Size**: Slices are more flexible than arrays and can grow or shrink.
- **Declaration**: `[]type{values}` or creating from an array `arr[low:high]`.
- **The `make()` function**: Creating slices with predefined length and capacity: `make([]type, length, capacity)`.
- **Memory Management**: Understanding how slices are abstractions over arrays.
- **Utility Functions**:
  - `len()`: Current number of elements.
  - `cap()`: Maximum capacity before re-allocation.
  - `copy()`: Efficiently copying elements from one slice to another.

## 📂 Files

- [main.go](main.go): The entry point for Day 5, coordinating array and slice demonstrations.
- [array.go](array.go): Comprehensive examples of array declaration, indexing, and iteration.
- [slice.go](slice.go): Detailed practice with slices, including `make()`, sub-slicing, and `copy()`.

## 🚀 How to Run

From the project root:
```bash
go run day5/main.go
```
