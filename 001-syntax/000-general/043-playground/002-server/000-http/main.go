package main

import (
	"fmt"
	"net/http"

)

func main() {
	// Create a new http.Server object.
	server := http.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write a message to the response writer.
		fmt.Fprintf(w, "Hello, world!")
	}))

	// Listen on port 8080.
	err := server.ListenAndServe(":8080")
	if err != nil {
		// Handle any errors.
		panic(err)
	}
}
