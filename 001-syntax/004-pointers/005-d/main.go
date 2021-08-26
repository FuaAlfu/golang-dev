package main

import "fmt"

func main() {
	i, j := 22, 33

	fmt.Println("value of i: [", i, "]\t\tand the address is: [", &i, "]")
	fmt.Println("value of j: [", j, "]\t\tand the address is: [", &j, "]")

	//p = pointer
	p := &i
	fmt.Println("value of p: [", *p, "]\t\tand the address is: [", p, "]")
	println("---")

	squarValue(j)
}

func squarValue(val int) {
	val *= val
	fmt.Println("value : [", val, "]\t\tand the address is: [", &val, "]")
}
