package main

import "fmt"

func main() {

	//create a slice
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	println("---")
	// 2# init a new elements
	x[0] = 33
	x[3] = 456121
	x[9] = 929
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	println("---")
	//3# appending
	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Print(cap(x))
	println("---")
	//4# Add Size
	x = append(x, 3424)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Print(cap(x))
}
