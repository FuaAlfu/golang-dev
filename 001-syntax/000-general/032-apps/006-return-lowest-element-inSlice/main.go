package main

import "fmt"

func findLowest(a []int) {
	smallestElement := a[0]

	for _, element := range a {
		if element < smallestElement {
			smallestElement = element
		}
	}
	fmt.Println("the lowest number of given slice is: ", smallestElement)
	fmt.Println("smallestElement:\tvalue of [", smallestElement, "]\taddress of[", &smallestElement, "]")
}

func main() {
	s := []int{42, 7, 1009, 78}
	findLowest(s)
}
