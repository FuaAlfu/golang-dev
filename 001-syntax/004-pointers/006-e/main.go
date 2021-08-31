package main

import "fmt"

func squaredAdd(p *int) {
	*p *= *p
	fmt.Println("value of p: [", *p, "]\t\tand the address is: [", p, "]")
}

func main() {
	v := 2

	squaredAdd(&v)
}
