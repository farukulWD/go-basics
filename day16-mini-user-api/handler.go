package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
		}

		u.ID = nextId
		nextId++
		users = append(users, u)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL: /users/3 → "3"
	// Extract ID from URL: /users/3 → "3"
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		for _, u := range users {
			if u.ID == id {
				json.NewEncoder(w).Encode(u)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)

	case http.MethodDelete:
		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return

			}
		}
		http.Error(w, "User not found", http.StatusNotFound)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
