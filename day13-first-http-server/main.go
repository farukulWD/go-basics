package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string `json:"name"`
	Grade string `json:"grade"`
}

var students = []Student{
	{"Rahim", "A"},
	{"Sadia", "B"},
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var s Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	students = append(students, s)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

func main() {
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getStudents(w, r)
		case http.MethodPost:
			addStudent(w, r)
		default:
			http.Error(w, "not allowed", 405)
		}
	})
	fmt.Println("API running on :8080")
	http.ListenAndServe(":8080", nil)
}
