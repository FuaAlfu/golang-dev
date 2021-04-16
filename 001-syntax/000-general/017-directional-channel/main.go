package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	fmt.Println("---")
	//show chan type
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
	println("---")

	//General to specific assigns ..
	cr = c
	cs = c

	//specific to general doesn't convert
	fmt.Print("====")
	//fmt.Printf("c\t%T\n", (chan int)(cr)) //not work
	//fmt.Printf("c\t%T\n", (chan int)(cs)) //not work
}
