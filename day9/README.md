# Day 9: Practical Projects in Go

This directory features small, real-world mini-apps that combine all the core concepts learned so far: variables, control flow, functions, structs, and maps.

## 📝 Projects Included

### 📊 Student Grade Analyzer
- **Location**: [main.go](main.go) and [grades.go](grades.go)
- **Features**: Uses a Map to store student scores and a function with conditional logic to calculate letter grades.

### 📔 Contact Book CLI
- **Location**: [contact_book/main.go](contact_book/main.go)
- **Features**: An interactive terminal app that uses structs and maps to add, search, and list contacts. Demonstrates infinite loops (`for { }`), `switch`, and user input using `fmt.Scan`.

### 🧮 Word Counter (Planned/Demo)
- **Location**: [world_counter/main.go](world_counter/main.go)
- **Features**: Designed to count word frequency in a given string using maps.

### 🎓 Student Profiles
- **Location**: [student_profiles/](student_profiles/)
- **Features**: Detailed student management using complex data structures.

## 🚀 How to Run

Navigate into any project directory and run:
```bash
go run .
```

Example for Student Grades:
```bash
go run day9/main.go day9/grades.go
```

Example for Contact Book:
```bash
go run day9/contact_book/main.go
```
