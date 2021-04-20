package main

import (
	"fmt"
)

var s string
var plus int

func main() {
	sliceString := []string{"a", "b", "c"}
	sliceInt := []int{11, 33, 55}

	rangeIntSlices(sliceInt)
	rangeStringSlices(sliceString)
}

func rangeIntSlices(arg []int) {
	for i := 0; i < 5; i++ {
		fmt.Println("slice", append(arg))
		fmt.Println("go: ", i, arg[i])
	}
	println("!=end=!")
}

func rangeStringSlices(arg []string) {
	for i := 0; i < 5; i++ {
		fmt.Println("slice", append(arg))
		fmt.Println("go: ", i, arg[i])
	}
	println("!=end=!")
}
