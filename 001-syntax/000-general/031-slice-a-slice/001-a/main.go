package main

import "fmt"

func main() {
	//a slice
	s := []int{11, 22, 33, 44}
	ss := strings.Split(s)

	//show the second element of slice
	fmt.Println(s[1])

	//slice [from : Go To]
	fmt.Println(s[1:])

	//slice from,.. to..
	fmt.Println(s[1:3])

}
