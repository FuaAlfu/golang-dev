package main

import (
	"fmt"

)

type (
	person struct {
		Name string
		Age  int
	}
)

var gold int

func main() {
	p := person{
		Name: "fua",
		Age:  30,
	}

	plat := 300
	var gold *int = &plat
	silver := &plat
	fmt.Println("Hello World")
	println("hi World")
	println("---")
	fmt.Println("your name is [", p.Name, "] and\t your age is about [", p.Age, "]")
	println("---")
	fmt.Println("gold: ", gold)
	fmt.Println("silver: ", silver)
	fmt.Println("silver val: ", *silver)
}
