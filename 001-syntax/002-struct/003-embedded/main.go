package main

import "fmt"

type person struct {
	first string
	age   int
}
type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	sa1 := secretAgent{person: person{
		first: "fua",
		age:   30,
	},
		ltk: true,
	}
	p2 := person{
		first: "jhon",
		age:   52,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.age)
	fmt.Println(p2.first, p2.age)

}
