package main

import "fmt"

//using iota
const (
	//without using Bit Shifting..
	//kb := 1024
	//mb := 1024 * kb
	//gb := 1024 * mb

	_  = iota             // _  not been used here..
	kb = 1 << (iota * 10) //using bit shifting..
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	//bit shifting..
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)
	//without using iota
	//kb := 1024
	//mb := 1024 * kb
	//gb := 1024 * mb
	//
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)

}
