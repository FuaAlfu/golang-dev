package main

import (
	"fmt"
)

func main() {

	/*
		literal type or struct
		it is a type has no name
	*/

	//zero value
	var v1 struct {
		flag    bool
		counter int16
		pi      float32
	}
	fmt.Printf("%+v\n", v1)
	println("---")

	//add values
	v2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 1000,
		pi:      3.141592,
	}
	fmt.Printf("%+v\n", v2)
	fmt.Println("flag: [", v2.flag, "] and\t counter: [", v2.counter, "]")
	println("---")
}
