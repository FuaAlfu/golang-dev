package main

import(
	"fmt"
)

func findMax(a...int)int{
	min := a[0]
	for i := 0; i < len(a); i++{
		if a[i] > min{
			min = a[i]
		}
	}
	return min
}

func main() {
	s := []int{55,9,1,777,123,91,5,26,41}
	fmt.Printf("the mxn value is: %d",findMax(s...))
}