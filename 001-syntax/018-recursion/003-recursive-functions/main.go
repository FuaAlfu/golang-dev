package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

var fib func(n int) int

func main() {
	number := 5
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	//-------(printing)
	fmt.Println(fact(number))
	fmt.Println(fib(number))
	println("===")
	fmt.Printf("fact: %d \n", fact(number))
	fmt.Printf("fib: %v \n", fib(number))
	fmt.Printf("fact: %T \n", fact(number))
}
