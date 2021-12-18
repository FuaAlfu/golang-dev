package main

import (
	"fmt"
)

func getMiddleElement(s []int) int {
	mid := s[len(s)/2]
	return mid
}

func main() {
	s := []int{3, 88, 51, 999, 1, 12}

	fmt.Println(getMiddleElement(s))
}
