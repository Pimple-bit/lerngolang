package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	body, _ := ioutil.ReadAll(r.Body)
	var msg Message
	if err := json.Unmarshal(body, &msg); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	msg.ID = uuid.New().String() // генерируем уникальный ID

	mu.Lock()
	messages = append(messages, msg)
	mu.Unlock()

	fmt.Println("Добавлено сообщение:", msg)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message saved with ID: " + msg.ID + "\n"))
}


func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	idToDelete := string(body)

	mu.Lock()
	defer mu.Unlock()

	index := -1
	for i, msg := range messages {
		if msg.ID == idToDelete {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	removed := messages[index]
	messages = append(messages[:index], messages[index+1:]...)

	fmt.Println("Удалено сообщение:", removed)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted message with ID: " + removed.ID + "\n"))
}


func allMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
	fmt.Println("Список всех сообщений отправлен клиенту")
}

func main() {
	http.HandleFunc("/add", addMessageHandler)
	http.HandleFunc("/delete", deleteMessageHandler)
	http.HandleFunc("/all", allMessagesHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}