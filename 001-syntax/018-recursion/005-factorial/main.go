package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	x := 32
	fmt.Printf("fact of %d is %d ", x, factorial(x))
}
