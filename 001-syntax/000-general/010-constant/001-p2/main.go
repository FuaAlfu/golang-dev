package main

import "fmt"

const(
	//max integer :: value on 64 bit architecture
	maxInt = 776549999877909877

	//much larger  value than int64
	bigger = 776549999877909877999877

	/*
	will not compile
	biggerInt int64 = 776549999877909877999877
	*/
)

func main() {
	/*
	show how much of high level precision we have in that constant system
	and they are not realy variables at all
	*/
	fmt.Println("maxInt: [", maxInt, "] and\t bigger: [", bigger, "]")
	fmt.Printf("maxInt[%d] bigger[%d]\n ", maxInt, bigger)
	println("---")
}