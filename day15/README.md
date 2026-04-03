# Day 15: Mastering JSON in Go

This directory explores how to work with JSON data in Go, covering encoding (marshalling), decoding (unmarshalling), and advanced struct tags for API design.

## 📝 Key Concepts Covered

- **The `encoding/json` Package**: Using Go's built-in support for generating and parsing JSON.
- **Marshalling (Go → JSON)**:
  - `json.Marshal`: Converting structs into compact JSON bytes.
  - `json.MarshalIndent`: Creating "pretty" human-readable JSON output with indentation.
- **Unmarshalling (JSON → Go)**:
  - `json.Unmarshal`: Converting JSON string/bytes into Go structs.
- **Advanced Struct Tags**:
  - `json:"name"`: Renaming struct fields in the resulting JSON.
  - `json:",omitempty"`: Excluding fields from JSON when their value is empty.
  - `json:"-"`: Completely hiding sensitive fields (like passwords) from being serialized.
- **JSON Streams (HTTP & Files)**:
  - `json.NewEncoder`: Directly streaming JSON response back to an HTTP client.
  - `json.NewDecoder`: Efficiently parsing incoming JSON request bodies without reading the whole body at once.

## 📂 Files

- [main.go](main.go): Practical demonstration of all JSON operations including struct tags, marshalling, unmarshalling, and basic HTTP routing for JSON data.

## 🚀 How to Run

1.  Run the application from the project root:
    ```bash
    go run day15/main.go
    ```
2.  Test the JSON endpoints:
    - **GET JSON**: `curl http://localhost:8080/hello`
    - **POST JSON**: `curl -X POST -H "Content-Type: application/json" -d '{"name":"Karim","age":25}' http://localhost:8080/create`
