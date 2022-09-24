package main

import(
	"fmt"
)

func findMin(a...int)int{
	min := a[0]
	for i := 0; i < len(a); i++{
		if a[i] < min{
			min = a[i]
		}
	}
	return min
}

func main() {
	s := []int{55,9,1,777,123,91,5,26,41}
	fmt.Printf("the min value is: %d",findMin(s...))
}