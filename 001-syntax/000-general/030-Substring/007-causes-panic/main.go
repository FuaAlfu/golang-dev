package main

import "fmt"

func main() {
	value := "abcd"
	/*
		this code is broken..
		 Indexes must be validated first or a panic occurs.
	*/
	result := value[10:60]
	fmt.Println(result)
}
