package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequest()
	println("serving at localhost: 8000")
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	//reach endpoint
	fmt.Fprint(w, "access to main page")
}

func handleRequest() {
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
