package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	
	codeStr := r.URL.Query().Get("code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		http.Error(w, "Invalid code", http.StatusBadRequest)
		return
	}

	switch code {
	case 200:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Everything is OK!")
	case 400:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad request!")
	case 404:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Resource not found!")
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal server error!")
	default:
		w.WriteHeader(http.StatusTeapot) // 418 :-)
		fmt.Fprintln(w, "Unknown status code!")
	}
}

func main() {
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}