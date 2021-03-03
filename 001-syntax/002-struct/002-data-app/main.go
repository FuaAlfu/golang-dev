package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//later
}

func createHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	person := struct {
		name string `json:"name"`
		age  int    `json:"age"`
	}{}
	err := decoder.Decode(&person)
	if err != nil {
		log.Println(err)
		return
	}
	printInfo(person.name, person.age)
	return
}

func printInfo(a string, b int) {
	fmt.Println(a, b)
}
