package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
		Write and Read:
		http.SetCookie()
		req.Cookie()
		http.StatusText()
	*/
	http.HandleFunc("/", set)

	http.HandleFunc("/read", read)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, req *http.Request) {

	//here
	http.SetCookie(w, &http.Cookie{

		Name: "my-cookie",

		Value: "some value",

		Path: "/",
	})

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")

	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")

}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie") //here

	if err != nil {

		http.Error(w, http.StatusText(400), http.StatusBadRequest)

		return

	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)

}
