package main

import (
	"fmt"
	"net/http"
)

func backend2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from backend 2")
}

func main() {
	http.HandleFunc("/", backend2Handler)

	fmt.Println("Backend 2 listening on :9002")
	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
