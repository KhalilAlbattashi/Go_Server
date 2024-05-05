package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request URL path exactly matches "/hello"
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
	// Check if the request method is a GET
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Create a new file server handler
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Register the handler function for the "/hello" URL pattern
	http.HandleFunc("/hello", GreetingHandler)

	// Register the handler function for the "/form" URL pattern
	http.HandleFunc("/form", formHandler)

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
