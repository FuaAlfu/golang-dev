package main

import "fmt"

func main() {
	value := "fua-alfu"
	getsAllPossibleSubstrings(value)
}

func getsAllPossibleSubstrings(s string) {
	//Convert to rune slice for substrings.
	runes := []rune(s)

	/* Loop over possible lengths, and possible start indexes.
	... Then take each possible substring from the source string.
	*/
	for length := 1; length < len(runes); length++ {
		//nested loop
		for start := 0; start <= len(runes)-length; start++ {
			substring := string(runes[start : start+length])
			fmt.Println(substring)
		}
	}
}
