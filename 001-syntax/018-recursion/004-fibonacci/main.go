package main

import "fmt"

func fibonacci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

//zero value
var i int

func main() {
	x := 100
	for i = 0; i <= x; i++ {
		fmt.Printf("f: %d \n", fibonacci(i))
	}
}
