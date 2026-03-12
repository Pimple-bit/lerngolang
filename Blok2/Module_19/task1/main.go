package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Age       int     `json:"age"`
	Married   bool    `json:"married"`
	Height    float64 `json:"height"`
}


func readUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Получен пользователь:", user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User received successfully\n"))
}


func sendUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	user := User{
		Name:    "Alice",
		Address: "123 Main St",
		Age:     30,
		Married: true,
		Height:  1.68,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/read-user", readUserHandler)
	http.HandleFunc("/send-user", sendUserHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}