package main

import (
	"fmt"
)

//zero value
var plus int //Plus = 0

func main() {
	/*
		nested loop
	*/
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	fmt.Print("Type *\n")

	for i := 9; i < 11; i++ {
		for j := 5; j < 10; j++ {
			for k := 3; k < j; k++ {
				fmt.Print("*", plus)
			}
			fmt.Println("*")
		}
		fmt.Println("*")
	}
}
