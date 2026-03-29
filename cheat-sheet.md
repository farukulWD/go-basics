# Go Basics Cheat Sheet 📝

A quick reference for the core concepts covered in this project.

## 🏃 Execution

```bash
go run main.go          # Execute a file
go build main.go        # Build a binary
go fmt main.go          # Format code
```

## 📦 Package & Import

```go
package main

import (
    "fmt"
    "math"
)
```

## 🔢 Variables & Types

```go
// Declarations
var x int = 10
y := 20                 // Recommended (only inside functions)
const Pi = 3.14

// Common Types
int, float64, string, bool

## 🍱 Arrays & Slices

```go
// Arrays (Fixed size)
var arr = [3]int{1, 2, 3}
arr2 := [...]int{1, 2, 3, 4} // Length inferred

// Slices (Dynamic size)
slice := []int{10, 20}
slice = append(slice, 30) // Grow slice

// make() function
s := make([]int, 5, 10) // len=5, cap=10

// range iteration
for index, value := range slice {
    fmt.Println(index, value)
}
```
```

## 🎮 Control Flow

### If / Else
```go
if age > 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

### Switch
```go
switch day {
case "Friday":
    fmt.Println("Almost weekend!")
default:
    fmt.Println("Busy day")
}
```

### Loops
```go
// For
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While (represented by for)
for x < 10 {
    x++
}

// Range (for slices/maps)
for i, v := range slice {
    fmt.Println(i, v)
}
```

## 🛠️ Functions

```go
func add(a int, b int) int {
    return a + b
}

func swap(a, b string) (string, string) {
    return b, a // Multiple return values
}

## 🏗️ Structs & Methods

```go
type User struct {
    Name string
    Age  int
}

// Method (Receiver function)
func (u User) Greet() {
    fmt.Printf("Hello, I'm %s!\n", u.Name)
}
```

## 🗺️ Maps

```go
// Declaration
m := make(map[string]int)
m["Age"] = 30

// Literal
ages := map[string]int{"Alice": 25, "Bob": 35}

// Iteration
for k, v := range ages {
    fmt.Println(k, v)
}

// Delete key
delete(ages, "Alice")
```

## 📍 Pointers

```go
var x int = 10
ptr := &x      // & gets address
fmt.Println(ptr)  // memory address
fmt.Println(*ptr) // * dereferences (gets value)

*ptr = 20      // Change original value via pointer

// Function with pointer
func change(val *int) {
    *val = 100
}
```

## 🧩 Interfaces

```go
type Shape interface {
    Area() float64
}

// Any type with Area() method implements Shape
type Circle struct { Radius float64 }
func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

// Empty interface (holds any type)
var i any = "Hello"

// Type Switch
switch v := i.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
}
```
```

---
*Keep practicing!*
