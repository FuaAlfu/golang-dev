package main

import "fmt"

/*
type( 
	Car struct{
	Wheels int
	Doors int
	Color string
}

    Info interface{}
)

func (i *Info) CarInfo() *Car {
	return &Car{Info: i}
}
*/
func foo(number int) *int{
	return &number
}

func bar() *int{
	number := 42
	return &number
}

func main() {
	number := foo(32)
	number2 := bar()

	fmt.Printf("number: %d \n", *number)
	fmt.Printf("number2: %d", *number2)
}