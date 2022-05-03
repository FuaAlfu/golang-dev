package main

import (
	"encoding/json"
	"fmt"
    "log"

)

type Career struct {
	Title  string `json:"title"`
}

type Person struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Career Career  `json:"career"`
}

func main() {
	//constructer
	p1 := Person{
		ID:   67654330098,
		Name: "Jhon",
		Career: Career{
			Title:  "teacher",
		},
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
