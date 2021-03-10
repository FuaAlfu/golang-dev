package main

import "fmt"

func main() {
	a := 30
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Println("the valu of a is: ",a," and the pointer is: ", &a)
	println("---")

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Println(*&a)
	println("---")

	//print all c value
	c := 0
	fmt.Println("c befor", &c)
	fmt.Println("c befor", c)
	pointer(&c)
	fmt.Println("c after", &c)
	fmt.Println("c after", c)
}

func pointer(p *int) {
	fmt.Println("p befor", p)
	fmt.Println("p befor", *p)
	*p = 30
	fmt.Println("p after", p)
	fmt.Println("p after", *p)
}