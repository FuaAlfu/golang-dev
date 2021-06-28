package main

import (
	"fmt"
	"strings"

)

func main() {
	//a string
	s := "11, 22, 33, 44"
	ss := strings.Split(s, ",")
	fmt.Println(ss)
	fmt.Println("The length of the slice is: ", len(ss))
	fmt.Printf("%d ", ss)
}
