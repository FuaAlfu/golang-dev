package main

import(
	"fmt"
	_"github.com/stretchr/testify/assert"
)

type Person struct {
	name, age interface{}
}

func multipleValues(s string, n int) (string, int) {
	return s, n
}

func main() {
	//constructing..
	p1 := Person{"fua", 29}
	//p2 := Person{"mua", 24}
	p3 := multipleValues("john",56)
	fmt.Println("p1 name is: [", p1.name, "]\tand p1 age is[", p1.age, "]")
	fmt.Println("p1 name address is: [", &p1.name, "]\tand p1 age address of[", &p1.age, "]")
	println("---")
	//assert.Equal(multipleValues(p2.name,p2.age))
	//println(multipleValues(p2.name,p2.age))
	println(p3)

}