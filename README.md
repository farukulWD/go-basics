# Go Basics Learning Journey 🚀

Welcome to my Go programming language learning repository! This project documents my progress as I learn Go from the ground up, starting with core concepts and moving towards more advanced features.

## 📁 Project Structure

The project is organized by days of learning, each focusing on specific topics:

- **Day 1**: [Day 1: Hello World](day1/) - Setting up the environment and writing the first "Hello World" program.
- **Day 2**: [Day 2: Basics](day2/) - Introduction to basic syntax and structure.
- **Day 3**: [Day 3: Data Types](day3-data-types/) - Deep dive into Go's data types and variables.
- **Day 4**: [Day 4: Control Flow](day4-control-flow/) - Mastering if/else, switch, and loops.

## 🛠️ Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.26.1 or later)
- `make` utility

### How to Run

To make development easier, I've included a `Makefile` with several helpful targets:

1.  **List all programs**:
    ```bash
    make list
    ```

2.  **Run a specific program**:
    ```bash
    make run FILE=day1/helloworld.go
    ```

3.  **Build a binary**:
    ```bash
    make build FILE=day1/helloworld.go
    ```

4.  **Clean build artifacts**:
    ```bash
    make clean
    ```

## 📝 Learning Notes

My goal is to understand Go's unique approach to:
- Concurrency (Goroutines and Channels)
- Interfaces and Composition
- Error Handling
- Performance and Efficiency

---
*Follow along as I continue my journey with Go!*
