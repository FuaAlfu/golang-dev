package main

import "fmt"

func removeLastElement(a ...int) []int {
	s := a[:len(a)-1]
	return s
}

func main() {
	s := []int{3, 88, 51, 999, 1, 12}
	fmt.Println(removeLastElement(s...))
}
