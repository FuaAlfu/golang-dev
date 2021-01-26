package main

import "fmt"

type person struct {
	name string
	age  int16
	id   float32
}

func main() {
	//constructing...
	//zero value
	var p1 person

	//display the value
	fmt.Printf("%+v", p1)

	//empty literal
	//e := person{} or return person{}

	//a struct literal
	p2 := person{
		name: "fua",
		age:  30,
		id:   227.828982,
	}
	//display the value
	fmt.Printf("%#v", p2.name)
}
