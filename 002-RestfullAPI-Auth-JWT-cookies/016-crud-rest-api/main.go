package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"

)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUsers).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("POST")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("Hello")
	initialMigration()
	initializeRouter()
}
