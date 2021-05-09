package main

import "fmt"

type person struct {
	Name     string
	Age      int
	synonyms []string
}

func main() {
	p1 := person{
		Name: "foo",
		Age:  32,
		synonyms: []string{
			"I'm tired",
			"what were they thinking !!",
		},
	}

	fmt.Println(p1.synonyms[1])
}
