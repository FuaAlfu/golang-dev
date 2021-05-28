package main

import "fmt"

func main() {
	//printing the lovely ascii char..
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
