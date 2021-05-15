package main

import (
	"fmt"
	"log"
	"net/http"

)

func main() {
	handle()
	println("servering on port: 8080")
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)
}

func handle() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
