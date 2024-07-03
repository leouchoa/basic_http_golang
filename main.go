package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var PORT = 8080

var data = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func getData(w http.ResponseWriter, req *http.Request) {
	// Sets up the Content-Type header so that the client knows to expect a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/get_data", getData)

	fmt.Println("Server is listening on port " + strconv.Itoa(PORT))

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
