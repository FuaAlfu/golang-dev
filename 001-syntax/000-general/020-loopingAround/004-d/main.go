package main

import "fmt"

func main() {
	//print even numbers
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	println("---")

	//print odd numbers
	for j := 0; j < 100; j++ {
		if j%2 != 0 {
			fmt.Println(j)
		}
	}
}
