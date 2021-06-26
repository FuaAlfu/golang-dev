package main

import "fmt"

func main() {
	/*
		              -- handles Unicode characters --
			value: contains a string that represent Unicode characters.
			runes: Convert string to rune slice before taking substrings
			no needed for ASCII strings.
	*/
	value := "fa:ma"

	runes := []rune(value)
	fmt.Println("First: ", string(runes[0:1]))
	fmt.Println("rest: ", string(runes[2:]))
}
