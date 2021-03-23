package main

import "fmt"

func main() {

	//untyped constants
	const ui = 12345
	const uf = 3.141592

	//typed constants
	const ti int = 12345
	const tf float64 = 3.141592

	//const myUnit8 uint8 = 1000

	//constant third will be of kind floating point
	const third = 1 / 3.0

	//constant zero will be of kind integer
	const zero = 1 / 3

	/*
		this is an example of constant arithemetic between typed and
		untyped constant.
		must have like types to perform math.
	*/
	const one int8 = 1

	fmt.Println("ui: [", ui, "] and\t uf: [", uf, "]")
	fmt.Println("ti: [", ti, "] and\t tf: [", tf, "]")
	fmt.Printf("ui[%d] uf[%d]\n ", ui, uf)
	println("---")
	fmt.Println("third: [", third, "] and\t zero: [", zero, "]")
	fmt.Printf("third[%d] zero[%d]\n ", third, zero)

}
