package main

import "fmt"

func main() {
	/*
		this is how to code while in Go..
		not recommended
	*/
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Print("done..\n")
	println("----")

	/*
		this is how to code while in Go..
		with if statement..
	*/
	xi := 1
	for {
		if xi > 9 {
			break
		}
		fmt.Println(xi)
		xi++
	}
	fmt.Print("done..")
}
