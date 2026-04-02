package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
	Score    int    `json:"score,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	s := User{Name: "Rahim", Age: 21}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// Decoder → read from HTTP request body
func createHandler(w http.ResponseWriter, r *http.Request) {
	var s User
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "bad JSON", 400)
		return
	}
	fmt.Println(s.Name)
}

func main() {

	u := User{Name: "Rahim", Age: 22, Email: "example@gmail.com", Password: "abcjanina", Score: 99}

	// Marshal JSON
	// compact
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
	// {"name":"Rahim","age":22,"email":"example@gmail.com","score":99}

	// Pretty printed
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(pretty))
	/* {
	  "name": "Rahim",
	  "age": 22,
	  "email": "example@gmail.com",
	  "score": 99
	}*/
	// Unmarshal: JSON

	jsonData := []byte(`{"name":"Rahim","age":22,"email":"example@gmail.com","score":99}`)
	var s User
	err := json.Unmarshal(jsonData, &s)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(s.Name) // Rahim

	// Encoder and Decoder (for HTTP / files)

	http.HandleFunc("/hello", handler)
	http.HandleFunc("/create", createHandler)

	fmt.Println("API running on :8080")
	http.ListenAndServe(":8080", nil)

}
