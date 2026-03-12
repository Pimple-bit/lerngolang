package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

type Message struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	ZipCode int    `json:"zipCode"`
	Text    string `json:"text"`
	Urgent  bool   `json:"urgent"`
}

var (
	messages []Message
	mu       sync.Mutex
)


func addMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	msg.ID = uuid.New().String()

	mu.Lock()
	messages = append(messages, msg)
	mu.Unlock()

	fmt.Println("Добавлено сообщение:", msg)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message saved with ID: " + msg.ID + "\n"))
}


func allMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	filtered := messages


	zipParam := r.URL.Query().Get("zipCode")
	if zipParam != "" {
		zipInt, err := strconv.Atoi(zipParam)
		if err == nil {
			tmp := []Message{}
			for _, msg := range filtered {
				if msg.ZipCode == zipInt {
					tmp = append(tmp, msg)
				}
			}
			filtered = tmp
		}
	}

	
	urgentParam := r.URL.Query().Get("urgent")
	if urgentParam != "" {
		urgentBool, err := strconv.ParseBool(urgentParam)
		if err == nil {
			tmp := []Message{}
			for _, msg := range filtered {
				if msg.Urgent == urgentBool {
					tmp = append(tmp, msg)
				}
			}
			filtered = tmp
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
	fmt.Println("Отправлен список сообщений с фильтром:", r.URL.RawQuery)
}

func main() {
	http.HandleFunc("/add", addMessageHandler)
	http.HandleFunc("/all", allMessagesHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}