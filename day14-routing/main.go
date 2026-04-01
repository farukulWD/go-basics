package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ─── Model ────────────────────────────────────────────────────────────────────

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

// ─── "Database" ───────────────────────────────────────────────────────────────

type BookStore struct {
	books  []Book
	nextID int
}

func NewBookStore() *BookStore {
	return &BookStore{
		nextID: 4,
		books: []Book{
			{1, "The Go Programming Language", "Alan Donovan", "tech"},
			{2, "Clean Code", "Robert Martin", "tech"},
			{3, "Atomic Habits", "James Clear", "self-help"},
		},
	}
}

func (store *BookStore) findByID(id int) (Book, int, bool) {
	for i, b := range store.books {
		if b.ID == id {
			return b, i, true
		}
	}
	return Book{}, -1, false
}

// ─── Struct-based Handler (implements http.Handler) ───────────────────────────
//
// Instead of bare functions, we attach methods to BookStore.
// This lets handlers share state (the books slice) without globals.

// ServeHTTP makes *BookStore satisfy the http.Handler interface.
// It acts as the top-level router — decides which method to call.
func (store *BookStore) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Trim prefix and split path into segments
	// e.g. "/books/2" → ["books", "2"]
	path := strings.Trim(req.URL.Path, "/")
	segments := strings.Split(path, "/")

	switch {
	case len(segments) == 1 && segments[0] == "books":
		store.handleCollection(res, req)

	case len(segments) == 2 && segments[0] == "books":
		id, err := strconv.Atoi(segments[1])
		if err != nil {
			http.Error(res, "id must be a number", http.StatusBadRequest)
			return
		}
		store.handleResource(res, req, id)

	default:
		http.NotFound(res, req)
	}
}

// handleCollection → GET /books  |  POST /books
func (store *BookStore) handleCollection(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		genre := req.URL.Query().Get("genre")
		if genre == "" {
			respond(res, http.StatusOK, store.books)
			return
		}
		var result []Book
		for _, b := range store.books {
			if b.Genre == genre {
				result = append(result, b)
			}
		}
		respond(res, http.StatusOK, result)

	case http.MethodPost:
		var b Book
		if err := json.NewDecoder(req.Body).Decode(&b); err != nil {
			http.Error(res, "invalid JSON", http.StatusBadRequest)
			return
		}
		b.ID = store.nextID
		store.nextID++
		store.books = append(store.books, b)
		respond(res, http.StatusCreated, b)

	default:
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleResource → GET /books/:id  |  PUT /books/:id  |  DELETE /books/:id
func (store *BookStore) handleResource(res http.ResponseWriter, req *http.Request, id int) {
	switch req.Method {
	case http.MethodGet:
		b, _, ok := store.findByID(id)
		if !ok {
			http.Error(res, "book not found", http.StatusNotFound)
			return
		}
		respond(res, http.StatusOK, b)

	case http.MethodPut:
		_, idx, ok := store.findByID(id)
		if !ok {
			http.Error(res, "book not found", http.StatusNotFound)
			return
		}
		var updated Book
		if err := json.NewDecoder(req.Body).Decode(&updated); err != nil {
			http.Error(res, "invalid JSON", http.StatusBadRequest)
			return
		}
		updated.ID = id
		store.books[idx] = updated
		respond(res, http.StatusOK, updated)

	case http.MethodDelete:
		_, idx, ok := store.findByID(id)
		if !ok {
			http.Error(res, "book not found", http.StatusNotFound)
			return
		}
		store.books = append(store.books[:idx], store.books[idx+1:]...)
		res.WriteHeader(http.StatusNoContent)

	default:
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// ─── Middleware ───────────────────────────────────────────────────────────────
//
// Middleware wraps an http.Handler and returns a new http.Handler.
// This is the standard pattern — chain as many as you need.

// withLogger logs method, path, and response time.
func withLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(res, req)
		fmt.Printf("[%s] %-20s %s\n", req.Method, req.URL.Path, time.Since(start))
	})
}

// withJSON forces Content-Type: application/json on every response.
func withJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

// chain applies middlewares right-to-left so the first one runs first.
func chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// ─── Helper ───────────────────────────────────────────────────────────────────

func respond(res http.ResponseWriter, status int, data any) {
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(data)
}

// ─── Main ─────────────────────────────────────────────────────────────────────

func main() {
	store := NewBookStore()

	mux := http.NewServeMux()

	// Register the struct handler under /books/
	// chain() applies: withLogger → withJSON → store.ServeHTTP
	mux.Handle("/books", chain(store, withLogger, withJSON))
	mux.Handle("/books/", chain(store, withLogger, withJSON))

	// A plain HandleFunc for a simple route (no struct needed)
	mux.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(map[string]string{"status": "ok"})
	})

	fmt.Println("BookStore API running on :8080")
	fmt.Println()
	fmt.Println("Routes:")
	fmt.Println("  GET    /books                   → all books")
	fmt.Println("  GET    /books?genre=tech         → filter by genre")
	fmt.Println("  POST   /books                   → create book")
	fmt.Println("  GET    /books/1                 → get by id")
	fmt.Println("  PUT    /books/1                 → update by id")
	fmt.Println("  DELETE /books/1                 → delete by id")
	fmt.Println("  GET    /health                  → health check")

	http.ListenAndServe(":8080", mux)
}
