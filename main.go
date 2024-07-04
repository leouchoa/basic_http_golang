package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var PORT = 8080
var PORT_STR = strconv.Itoa(PORT)

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
}

func getData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func postData(w http.ResponseWriter, req *http.Request) {
	productId := req.PathValue("id")
	fmt.Fprintf(w, "created properties for product %s", productId)
}

func displaySingleItem(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	fmt.Fprintf(w, "displaying properties for product %s", productId)
}

func main() {
	// https://codewithflash.com/advanced-routing-with-go-122
	v0Router := http.NewServeMux()

	v0Router.HandleFunc("/", helloHandler)
	v0Router.HandleFunc("GET /get_data/{id}", getData)
	v0Router.HandleFunc("GET /get_product/{id}", displaySingleItem)
	v0Router.HandleFunc("POST /post_product/{id}", postData)

	router := http.NewServeMux()
	router.Handle("/v0/", http.StripPrefix("/v0", v0Router))

	fmt.Println("Server is listening on port " + PORT_STR)

	err := http.ListenAndServe(":"+PORT_STR, router)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
