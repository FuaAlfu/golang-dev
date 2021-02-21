package main

import "fmt"

func main() {
	/*
		for init; condition; post {}
	*/

	//for i := 0; i <= 10; i++{   //Outer loop
	//  for j := 0; j < 3; j++{    //Inner loop
	//  	fmt.Printf("The outer loop: %d\t The inner loop: %d\n ",i,j)

	//Another way to do it

	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n ", j)
		}

	}
}
