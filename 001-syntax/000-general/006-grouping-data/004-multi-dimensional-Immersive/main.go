package main

import "fmt"

func multi(a []int, b ...int) {
	ab := [][]int{a, b}
	fmt.Printf("ab:\tvalue of %T \n [", ab, "] and address of [", &ab, "]")
}

func main() {
	x := []int{1, 2, 4, 5, 8}
	y := []int{22, 24, 28, 32}

	xy := append(x, y...)

	fmt.Printf("x:\tvalue of %T \n [", x, "] and address of [", &x, "]")
	println("---")
	fmt.Printf("y:\tvalue of %T \n [", y, "] and address of [", &y, "]")
	println("---")
	fmt.Printf("xy:\tvalue of %T\n [", xy, "] and address of [", &xy, "]")
	println("---")
	multi(x, y...)
}
