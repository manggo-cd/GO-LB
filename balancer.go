package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// --- types from previous step ---

type Backend struct {
	URL *url.URL
}

type RoundRobinBalancer struct {
	backends []*Backend
	current  int
}

func (b *RoundRobinBalancer) Next() *Backend {
	if len(b.backends) == 0 {
		return nil
	}
	backend := b.backends[b.current]
	b.current = (b.current + 1) % len(b.backends)
	return backend
}

// --- main and handler ---

func main() {
	// Parse backend URLs
	backendURLs := []string{
		"http://localhost:9001",
		"http://localhost:9002",
	}

	var backends []*Backend

	for _, rawURL := range backendURLs {
		parsed, err := url.Parse(rawURL)
		if err != nil {
			panic(err)
		}
		backends = append(backends, &Backend{URL: parsed})
	}

	balancer := &RoundRobinBalancer{
		backends: backends,
		current:  0,
	}

	// Handler that *just* shows which backend would be picked.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		backend := balancer.Next()
		if backend == nil {
			http.Error(w, "no backends available", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Would forward to backend: %s\n", backend.URL.String())
	})

	fmt.Println("Load balancer listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting load balancer:", err)
	}
}
