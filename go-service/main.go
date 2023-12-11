package main

import (
	"net/http"
)

func setupAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Handle only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(""))
}

func main() {
	// Handle POST requests to the "/setup-account" endpoint
	http.HandleFunc("/setup-account", setupAccountHandler)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
