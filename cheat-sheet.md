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
```

---
*Keep practicing!*
