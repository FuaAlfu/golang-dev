package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/home", home)
	http.HandleFunc("/refresh", refresh)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
