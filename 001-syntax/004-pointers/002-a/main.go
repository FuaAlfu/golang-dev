package main

import "fmt"

func main() {
	//an example of value semantic..
	count := 10

	fmt.Println("count:\tvalue of [", count, "]\taddress of[", &count, "]")
	fmt.Println()

	//pass the value of the count
	increment(count)

	println("count:\tvalue of [", count, "]\taddress of[", &count, "]")
}

func increment(inc int) {
	//increment the value of inc
	inc++
	println("inc:\tvalue of [", inc, "]\taddress of[", &inc, "]")
}
