package main

import(
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
	  count rune numbers in string
	*/
	s := "GO is awesome ! ❤️😊❤️"
	n := utf8.RuneCountInString(s)

	fmt.Println(len(s))
	fmt.Printf("here is %v", n)
}