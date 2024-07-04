package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var PORT = 8080

type user_data struct {
	id        uint
	name      string
	role      string
	email     string
	phone     string
	contacted bool
}

var data = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
	w.WriteHeader(http.StatusOK)
}

func getData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func displaySingleItem(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	fmt.Fprintf(w, "displaying properties for product %s", productId)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/get_data/{id}", getData)
	mux.HandleFunc("/product/{id}", displaySingleItem)

	fmt.Println("Server is listening on port " + strconv.Itoa(PORT))

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), mux)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
