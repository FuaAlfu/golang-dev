package main

import "fmt"

func main() {
	goo()
	fmt.Println(myInt(3333))

	//default Anonymous func
	func() {
		fmt.Println("Anonymous func run")
	}()

	// Anonymous func pass an arg
	func(x int) {
		fmt.Println("the meaning of life: ", x)
	}(33)

	//Anonymous func pass two arg
	func(x int, y int) {
		fmt.Println("here is the sum of X and Y:", x+y)
	}(99, 1)

}

//Method by identifier
func goo() {
	fmt.Println("goo is, runing")
}

//return func (type)
func myInt(x int) int {
	return x
}
