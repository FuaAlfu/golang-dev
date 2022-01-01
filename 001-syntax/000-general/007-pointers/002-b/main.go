package main

import "fmt"

func mod(a *[5]int) {
	(*a)[0] = 120
}

func main() {
	s := [5]int{22, 99, 5, 342, 1}
	mod(&s)
	fmt.Println("slice:\tvalue of [", s, "]\taddress of[", &s, "]")
}
