package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	x := 15
	fmt.Println("x:\tvalue of [", x, "]\taddress of[", &x, "]")

	//constructing
	p1 := Person{"john", 32}

	//pointer to the struct
	p2 := &p1
	fmt.Println(p2)

	//updating the value of name and age
	p2.name = "doe"
	p2.age = 42
	fmt.Println(p2)

	fmt.Println("p1:\tvalue of [", p1, "]\taddress of[", &p1, "]")
	fmt.Println("p2:\tvalue of [", p2, "]\taddress of[", &p2, "]")

}
