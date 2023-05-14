package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new mux router.
	r := mux.NewRouter()

	// Handle the "/" route.
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	// Handle the "/about" route.
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the about page.")
	})

	// Listen on port 8080.
	http.ListenAndServe(":8080", r)
}
