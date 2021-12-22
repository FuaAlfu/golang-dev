package main

import "fmt"

type Slice struct {
	s []int
}

func reverseOne(s ...int) []int {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func reverseTwo(s []int) []int {
	newElement := make([]int, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		newElement = append(newElement, s[i])
	}
	return s
}

func reversThree(s ...int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	slice := Slice{
		s: []int{11, 22, 33, 44, 55, 66, 77},
	}

	fmt.Printf("%d\n", reverseOne(slice.s...))
	println("===")
	fmt.Printf("%v\n", reverseTwo(slice.s))
	println("===")
	fmt.Printf("%s\n", reversThree(slice.s...))
}
