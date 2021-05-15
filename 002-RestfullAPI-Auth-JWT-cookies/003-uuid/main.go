package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

)

func main() {
	/*
		    go get
		    github.com/satori/go.uuid

			to look for latest version
			go mod tidy
	*/

	http.HandleFunc("/", index)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session")

	if err != nil {

		id, _ := uuid.NewV4() //id := uuid.NewV4()

		cookie = &http.Cookie{

			Name: "session",

			Value: id.String(),

			// Secure: true,

			HttpOnly: true,

			Path: "/",
		}

		http.SetCookie(w, cookie)

	}

	fmt.Println(cookie)

}
