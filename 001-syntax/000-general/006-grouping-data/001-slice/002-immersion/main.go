package main

import (
	"fmt"
)

func main() {
	s := []int{123, 321}
	ss := []int{456, 654}

	sum(s, ss...)
}

func sum(y []int, x ...int) {
	z := append(y, x...)
	fmt.Println(z)
	println("---")
	for i, v := range z {
		println("index:", i, " value:", v)
	}
	println("---")
}
