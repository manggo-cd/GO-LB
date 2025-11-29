package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run backend_server.go <port>")
		fmt.Println("Example: go run backend_server.go 8081")
		os.Exit(1)
	}

	port := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("Response from backend server on port %s\n", port)
		log.Printf("Received request on port %s: %s %s", port, r.Method, r.URL.Path)
		fmt.Fprint(w, message)
	})

	log.Printf("Backend server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

