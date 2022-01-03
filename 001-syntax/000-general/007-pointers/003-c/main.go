package main

import "fmt"

func modify(s ...int) {
	s[0] = s[0] - 1
	fmt.Println("slice:\tvalue of [", s, "]\taddress of[", &s, "]")
}

func main() {
	s := [...]int{1, 55, 2, 9}
	inc := &s
	//inc++  //go doesn't support pointer arithmetic..
	fmt.Println("slice:\tvalue of [", inc, "]\taddress of[", &inc, "]")
	println("---")

	s2 := []int{1, 55, 2, 9}
	modify(s2...)
}
