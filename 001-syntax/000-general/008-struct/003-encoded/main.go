package main

import (
	"encoding/json"
	"log"
	"net/http"

)

type (
	Data struct {
		Person Person `json:"person"`
	}

	Response struct {
		Data Data `json:"data"`
	}

	Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

func connect(w http.ResponseWriter, r *http.Request) {

	data := Response{
		Data: Data{
			Person: Person{
				Name: "John Doe",
				Age:  32,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func handle() {
	http.HandleFunc("/", connect)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handle()
}
