package main

import "fmt"

func fibonacci(i int) int {
	if i < 2 {
		return i
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	fmt.Println(fibonacci(10))
}
