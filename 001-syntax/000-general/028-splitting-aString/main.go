package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		splitting a string into a slice
	*/

	groupOfThings := "tv,phone,pc,car"
	things := strings.Split(groupOfThings, ",")
	fmt.Println(things)
}
