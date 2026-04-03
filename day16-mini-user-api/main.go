package main

import (
	"fmt"
	"net/http"
)

var users = []User{}
var nextId = 1

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userByIDHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
