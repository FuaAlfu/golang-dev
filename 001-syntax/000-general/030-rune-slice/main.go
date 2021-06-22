package main

import (
	"fmt"

)

func main() {

	value := "GOLANG"

	//first way
	//handles any kind of rune in the string.
	runes := []rune(value)

	//Convert back into a string from rune slice.
	safeSubstring := string(runes[1:3])
	fmt.Println(" RUNE SUBSTRING: ", safeSubstring)

	// the second way
	//handles only ASCII correctly.
	asciiSubstring := value[1:3]
	fmt.Println("ASCII SUBSTRING:", asciiSubstring)
}
