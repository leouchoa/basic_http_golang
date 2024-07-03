package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", req.URL.Path)
}

func main() {
	// Define a simple handler function

	// Register the handler function for a specific route pattern
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
