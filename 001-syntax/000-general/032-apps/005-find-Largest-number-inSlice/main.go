package main

import "fmt"

func findLargest(a ...int) {
	var largestNumber, tempNum int

	for _, element := range a {
		if element > tempNum {
			tempNum = element
			largestNumber = tempNum
		}
	}
	fmt.Println("the largest number of given slice is: ", largestNumber)
}

func main() {
	s := []int{42, 7, 1009, 78}
	findLargest(s...)
}
