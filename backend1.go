package main

import (
	"fmt"
	"net/http"
)

func backend1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from backend 1")
}

func main() {
	http.HandleFunc("/", backend1Handler)

	fmt.Println("Backend 1 listening on :9001")
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
