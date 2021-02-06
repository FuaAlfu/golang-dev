package main

import "fmt"

type foo struct {
	//memory layout
	flag    bool
	counter int16
	pi      float32
}

func main() {
	fmt.Println("----")

	//declare a variable of an anonymous type and init :: using a struct literal
	anonymous := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 100,
		pi:      3.141592,
	}

	//create a value of type foo :: zero value
	var f foo

	//assign the value of the unnamed struct type
	//to the named struct type value..
	f = anonymous

	//print the values
	fmt.Printf("%+v\n", f)
	fmt.Printf("%+v\n", anonymous)
	fmt.Println("flag", anonymous.flag)
}
