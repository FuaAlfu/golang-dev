package main

import (
	"fmt"
)

func main() {
	s := "smogy"
	ss := "mogy"
	x := "potion"
	xs := "runing !!" //here
	if x == "potion" {
		fmt.Println(healing(s))
	}

	if xs == "runing" {
		fmt.Println(runing(ss))
	}
}

func runing(s string) string {
	return s + " is runing !!"
}

func healing(s string) string {
	return s + " is healing !!"
}
