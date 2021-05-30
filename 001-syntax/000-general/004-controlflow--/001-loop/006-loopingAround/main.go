package main

import (
	"fmt"

)

var gold int

func ranging(s []int) {
	for _, sr := range s {
		fmt.Println("nums:\t [", sr, "]")
	}
}

func main() {
	g := 33
	fmt.Println(gold)
	fmt.Println(g)
	//----
	s1 := []int{11, 44, 77}
	ranging(s1)
	println("---")

	//-----------
	s2 := []string{"f", "g", "h"}
	for i := 0; i < len(s2); i++ {
		fmt.Println("Val is:\t", s2[i])
	}

	for i, v := range s2 {
		fmt.Printf("the value at index %v is %v \n", i, v)
	}
}
