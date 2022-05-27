package main

import(
	"fmt"
)

type Person struct {
	name, age interface{}
}

func main() {
	//constructing..
	p1 := Person{"fua", 29}
	p2 := Person{"mua", 24}

	fmt.Println("p1 name is: [", p1.name, "]\tand p1 age is[", p1.age, "]")
	fmt.Println("p1 name address is: [", &p1.name, "]\tand p1 age address of[", &p1.age, "]")
	println("---")
	println(p2.name)

}