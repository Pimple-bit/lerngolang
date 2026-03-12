package main

import (
	"fmt"
	"log"
	"net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}


	name := r.Header.Get("X-My-Name")
	if name == "" {
		name = "Guest"
	}

	
	greeting := fmt.Sprintf("Hello, %s!", name)
	fmt.Println(greeting)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(greeting + "\n"))
}

func main() {
	http.HandleFunc("/headers", headersHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}