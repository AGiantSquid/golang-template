package main

import (
	"log"
	"net/http"

	"github.com/AGiantSquid/golang-template/internal/handler"
)

func main() {
	// Register our handler
	http.HandleFunc("/", handler.HelloHandler)

	// Start the server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
