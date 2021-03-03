package main

import (
	"encoding/json"
	"fmt"
)

type company struct {
	CompanyName string
	Town        string
	Active      bool
	Employees   []Employees
}

type Employees struct {
	Name string
	Age  int
	Rank string
	XP   []string
}

func main() {
	//slice byte
	input := []byte(`
    {
		"CompanyName": "sony",
		"town": "tokyo",
		"active": true,
		"employees": [
		  {
			"name": "jhon",
			"age": 42,
			"rank": "senior",
			"xp": [
			  "eat",
			  "sleep",
			  "walking around"
			]
		  },
		  {
			"name": "sarah",
			"age": 23,
			"rank": "junior",
			"xp": [
			  "work hard",
			  "working harder",
			  "no sleep"
			]
		  }
		]
	  }
`)

	s := company{}
	err := json.Unmarshal(input, &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", s)
}
