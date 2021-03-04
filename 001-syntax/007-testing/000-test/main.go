package main

import "fmt"

func main() {
	fmt.Println("10 + 9 =", sum(10, 9))

}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
