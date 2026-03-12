package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "sync"

    "github.com/google/uuid"
)

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var (
	messages []Message
	mu       sync.Mutex
)

func addMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	text := string(body)

	msg := Message{
		ID:   uuid.New().String(),
		Text: text,
	}

	mu.Lock()
	messages = append(messages, msg)
	mu.Unlock()

	fmt.Println("Added message:", msg)
	fmt.Println("All messages:", messages)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message saved with ID: " + msg.ID + "\n"))
}

func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
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

	fmt.Println("Deleted message:", removed)
	fmt.Println("All messages:", messages)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted message with ID: " + removed.ID + "\n"))
}

func main() {
	http.HandleFunc("/add", addMessageHandler)
	http.HandleFunc("/delete", deleteMessageHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}