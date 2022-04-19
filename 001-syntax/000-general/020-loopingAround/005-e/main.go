package main

import(
	"fmt"
)

var count int

func main() {
	/*
	  infinite loop
	*/
	limit := 50
	for true{
		count++
		fmt.Printf("counter: %d ||", count)
		if count > limit{
			break;
		}
	}
	
}