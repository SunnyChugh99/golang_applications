package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response struct to represent the JSON response
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Define a handler function for the API endpoint
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		// Create a response message
		response := Response{Message: "Hello, World!"}

		// Convert the response struct to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to indicate JSON content
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the response writer
		w.Write(jsonResponse)
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
