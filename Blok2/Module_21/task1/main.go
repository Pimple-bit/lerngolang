package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)


type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Pages       int       `json:"pages"`
	Read        bool      `json:"read"`
	AddedAt     time.Time `json:"addedAt"`
	FinishedAt  time.Time `json:"finishedAt,omitempty"`
}

var (
	library []Book
	mu      sync.Mutex
)

// POST /books 
func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	book.ID = uuid.New().String()
	book.AddedAt = time.Now()

	mu.Lock()
	library = append(library, book)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	fmt.Println("Добавлена книга:", book)
}

// PATCH /books/id/read 
func markReadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/books/")
	id = strings.TrimSuffix(id, "/read")

	mu.Lock()
	defer mu.Unlock()

	for i, book := range library {
		if book.ID == id {
			if book.Read {
				http.Error(w, "Book already marked as read", http.StatusBadRequest)
				return
			}
			library[i].Read = true
			library[i].FinishedAt = time.Now()
			json.NewEncoder(w).Encode(library[i])
			fmt.Println("Книга прочитана:", library[i])
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// GET /books/
func getBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/books/")

	mu.Lock()
	defer mu.Unlock()

	for _, book := range library {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

// GET /books 
func listBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	filtered := library

	// Фильтр по автору
	author := r.URL.Query().Get("author")
	if author != "" {
		tmp := []Book{}
		for _, b := range filtered {
			if strings.EqualFold(b.Author, author) {
				tmp = append(tmp, b)
			}
		}
		filtered = tmp
	}

	// Фильтр по статусу прочтено
	readParam := r.URL.Query().Get("read")
	if readParam != "" {
		readBool, err := strconv.ParseBool(readParam)
		if err == nil {
			tmp := []Book{}
			for _, b := range filtered {
				if b.Read == readBool {
					tmp = append(tmp, b)
				}
			}
			filtered = tmp
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
	fmt.Println("Список книг отправлен с фильтром:", r.URL.RawQuery)
}

// DELETE /books/{id} 
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/books/")

	mu.Lock()
	defer mu.Unlock()

	index := -1
	for i, b := range library {
		if b.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	removed := library[index]
	library = append(library[:index], library[index+1:]...)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted book with ID: " + removed.ID + "\n"))
	fmt.Println("Книга удалена:", removed)
}

func main() {
	http.HandleFunc("/books", listBooksHandler)                // GET /books
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && !strings.HasSuffix(r.URL.Path, "/read"):
			getBookHandler(w, r)
		case r.Method == http.MethodPatch && strings.HasSuffix(r.URL.Path, "/read"):
			markReadHandler(w, r)
		case r.Method == http.MethodDelete:
			deleteBookHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/books/add", addBookHandler) // POST /books/add

	fmt.Println("Library REST API started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}