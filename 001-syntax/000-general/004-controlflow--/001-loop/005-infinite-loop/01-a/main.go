package main

import "fmt"

func main() {
	/*
		infinite loop in go
	*/

	for true {
		fmt.Println("This loop will keep runing forever..\n")
	}

	/*
		//or
		 i := 0
	    for {
	        fmt.Println(i)
	        i++
	    }
	*/
}
