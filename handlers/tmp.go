package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello Filip")
}

func main() {
	http.HandleFunc("/", hello)

	port := "10000"

	log.Printf("Server listening on port %s...", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
