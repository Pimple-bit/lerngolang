package main

import (
	"fmt"
	"log"
	"net/http"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем query параметры
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Printf("Получены query параметры: name=%s, age=%s\n", name, age)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello %s, age %s\n", name, age)))
}

func main() {
	http.HandleFunc("/query", queryHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}