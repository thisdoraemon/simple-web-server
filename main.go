package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}

func NewServer() http.Handler {
	fileServer := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/hello", helloHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" || r.Method != "GET" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing the form: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Request successfully reached\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
