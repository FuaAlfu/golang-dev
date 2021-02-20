package main

import "fmt"

const (
	a     = 20  //untyped constant...
	b int = 200 //Typed constant..

)

func main() {

	fmt.Println("The whole values is :\n", a, b)
	fmt.Printf("%T\n", a, b)

}
