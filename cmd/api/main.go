package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
