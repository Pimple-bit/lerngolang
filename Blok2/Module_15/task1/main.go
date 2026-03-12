package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var (
	messages []string
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

	mu.Lock()
	messages = append(messages, text)
	mu.Unlock()

	fmt.Println("Added message:", text)
	fmt.Println("All messages:", messages)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message saved\n"))
}

func main() {
	http.HandleFunc("/add", addMessageHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}