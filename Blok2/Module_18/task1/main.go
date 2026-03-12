package main

import (
	"fmt"
	"log"
	"net/http"
)

func postOnlyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Ошибка: неверный HTTP метод:", r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Хендлер успешно обработан методом POST")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST запрос успешно обработан!\n"))
}

func main() {
	http.HandleFunc("/post-only", postOnlyHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}