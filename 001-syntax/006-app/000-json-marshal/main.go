package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string `json:"name"`
	Age         int    `json:"age,omitempty"`
	Active      bool   `json:"-"`
	lastLoginAt string
}

func main() {
	/*
		json marshal struct with json tags
	*/
	u, err := json.Marshal(User{Name: "fua", Age: 30, Active: true, lastLoginAt: "today"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))

	/*output:
	{"full_name":"fua"}
	*/
}
