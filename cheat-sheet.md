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

## ⚠️ Error Handling

```go
// Basic handling
res, err := doSomething()
if err != nil {
    return fmt.Errorf("failed: %w", err) // %w wraps
}

// Sentinel Errors
var ErrNotFound = errors.New("not found")
if errors.Is(err, ErrNotFound) { ... }

// Custom Errors
type MyErr struct { Msg string }
func (e MyErr) Error() string { return e.Msg }

var me MyErr
if errors.As(err, &me) { ... }
```

## 🌐 HTTP Server

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

// Middleware Pattern
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // logic before
        next.ServeHTTP(w, r)
        // logic after
    })
}

// Custom Router (ServeMux)
mux := http.NewServeMux()
mux.Handle("/api/", middleware(myHandler))
```

## 📄 JSON

```go
type User struct {
    Name  string `json:"username"`    // rename
    Age   int    `json:",omitempty"`  // hide if empty
    Pass  string `json:"-"`          // ignore
}

// Marshalling (Struct to JSON)
data, _ := json.Marshal(user)
fmt.Println(string(data))

// Unmarshalling (JSON to Struct)
json.Unmarshal(data, &user)
```

## 🍸 Gin Framework

```go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.Run() // listens on :8080
}
```

## 🗄️ GORM & Databases

```go
import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

// Model with Tags
type User struct {
    gorm.Model           // contains ID, CreatedAt...
    Name string `gorm:"not null"`
}

// Connection
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// Operations
db.AutoMigrate(&User{})      // Sync DB schema
db.Create(&user)             // Insert
db.Find(&users)              // Fetch all
db.First(&user, id)          // Fetch by ID
db.Save(&user)               // Update
db.Delete(&user)             // Delete
```

## 🛡️ Validation (Gin)

```go
type CreateUser struct {
    Name  string `binding:"required,min=2"`
    Email string `binding:"required,email"`
    Age   int    `binding:"gte=18"`
}

// In controller
if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
}
```
```

---
*Keep practicing!*
