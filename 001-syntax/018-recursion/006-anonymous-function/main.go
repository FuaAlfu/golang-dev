package main

import "fmt"

var myFunction func()

func main() {
	myFunction = func() {
		fmt.Println("printing my function")
		myFunction()
	}

	myFunction()
}
