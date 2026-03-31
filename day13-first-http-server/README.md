# Day 13: First HTTP Server in Go

This directory marks the transition from CLI apps to the web, exploring how to build a basic REST API using Go's standard library.

## 📝 Key Concepts Covered

- **The `net/http` Package**: Using Go's powerful built-in package to handle web requests and server communication.
- **REST API Basics**:
  - **HTTP Methods**: Implementing logic for `GET` (fetching data) and `POST` (sending data).
  - **Routing**: Mapping URL paths to specific handler functions using `http.HandleFunc`.
- **JSON Handling**:
  - **Struct Tags**: Using tags like `` `json:"name"` `` to map Go struct fields to JSON property names.
  - **Encoding**: Converting Go slices/structs into JSON strings for the response (`json.NewEncoder`).
  - **Decoding**: Converting incoming JSON request bodies into Go structs (`json.NewDecoder`).
- **Server Lifecycle**: Starting the server on a specific port using `http.ListenAndServe(":8080", nil)`.

## 📂 Files

- [main.go](main.go): A mini-student management API that allows getting a list of students and adding new ones via JSON.

## 🚀 How to Run

1.  Run the server from the project root:
    ```bash
    go run day13-first-http-server/main.go
    ```
2.  Test the API using `curl` or a tool like Postman:
    - **GET (List Students)**:
      ```bash
      curl http://localhost:8080/students
      ```
    - **POST (Add Student)**:
      ```bash
      curl -X POST -H "Content-Type: application/json" -d '{"name":"Karim","grade":"A+"}' http://localhost:8080/students
      ```
