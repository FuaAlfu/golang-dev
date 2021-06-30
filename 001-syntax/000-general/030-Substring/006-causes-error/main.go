package main

import "fmt"

func main() {
	/*
			 This program does not work by badResult variable.
		    String slice cannot have three indexes.
	*/
	value := "one two three"

	//badResult not work
	//badResult := value[4:8:6]
	//fmt.Println(badResult)
	goodResult := value[5:8]
	anotherGoodResult := value[2:]
	fmt.Println(goodResult)
	fmt.Println(anotherGoodResult)
}
