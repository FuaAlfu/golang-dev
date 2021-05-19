package main

import (
	"fmt"
)

func main() {
	/*
		function is first class values
		litera means isn't named
	*/
	var n int

	func() {
		fmt.Println("Direct: ", n)
	}()
	//-----------------------

	//declare an anonymous function and assign it to a variable
	f := func() {
		fmt.Println("Variable: ", n)
	}

	//call the anonymous function through the variable
	f()

	//declare an anonymous function and assign value to a variable
	func(v int) {
		fmt.Println("Val: ", v)
	}(30)

	//defer the call to the anonymous function till after main returns.
	defer func() {
		fmt.Println("defer 1: ", n)
	}()

	//----------------------------------

	//set the value of n to 87 before the return
	n = 87

	//call the anonymous function through the variable :: once again
	f()

	//defer the call to the anonymous function till after main returns :: once again
	defer func() {
		fmt.Println("defer 2: ", n)
	}()
}
