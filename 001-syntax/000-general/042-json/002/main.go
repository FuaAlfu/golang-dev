package main

import (
	"encoding/json"
	"fmt"

)

type person struct {
	Name       string `"json:"name"`
	Age        int    `json: "age"` 
	City       string `json:"city"`
	RoomNumber int    `json:"number"`
}

var address map[string]interface{}

func main() {
	text := `{"name": "John", "age": 30, "city": "New York","number": 604}`
	textBytes := []byte(text)

	p1 := person{}
	err := json.Unmarshal(textBytes, &p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p1.RoomNumber)

	//more
	err = json.Unmarshal(textBytes, &address) 
	if err != nil {
		fmt.Println(err)
		return
	}
	println("---")
	for _, v := range address {
		fmt.Println(v)
	}
}
