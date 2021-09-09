package main

import "fmt"

type (
	person struct {
		name string
		age  int
	}
)

func initPerson() person {
	p := person{
		name: "john",
		age:  42,
	}
	return p
}

func initPersonDa() *person {
	p := person{
		name: "john",
		age:  42,
	}
	fmt.Printf("init person -> %p\n", &p)
	return &p
}

func main() {
	//return a copy of value in the active frame..
	fmt.Println(initPerson())
	println("---")

	//return a copy of an address in the active frame..
	fmt.Println(initPersonDa())
}
