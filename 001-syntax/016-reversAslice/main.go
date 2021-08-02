package main

import "fmt"

type sliceTrack struct {
	slice []int
}

func main() {
	ss := sliceTrack{
		slice: []int{1, 2, 3},
	}

	s := []int{11, 22, 33, 44, 55}

	fmt.Println(s)
	fmt.Printf("%v\n", reverseAslice(s...))
	fmt.Printf("%v\n", reverseAslice)
	println("---")
	fmt.Println(ss.slice)
	fmt.Printf("%v\n", reverseAslice(ss.slice...))
	fmt.Printf("%v\n", ss.slice)
}

func reverseAslice(numbers ...int) []int {
	newSlice := make([]int, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newSlice = append(newSlice, numbers[i])
	}
	return newSlice
}
