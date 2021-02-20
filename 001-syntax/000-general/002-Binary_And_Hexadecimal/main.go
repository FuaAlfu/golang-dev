package main

import "fmt"

func main() {

	/*
	   Binary And Hexadecimal
	*/

	s := "H"
	fmt.Print("\n", s)

	//[]byte called a slice byte
	bs := []byte(s)
	fmt.Print("\n", bs)
	println("-!-\n")

	n := bs[0] //Position
	fmt.Println(n)
	fmt.Printf("%v\n", n) //value
	fmt.Printf("%T\n", n) //Type
	fmt.Printf("%b\n", n) // Binary
	fmt.Printf("%#\n", n) //Hexadecimal
}
