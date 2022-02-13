package main

import "fmt"

func main() {
	
	/*
	 new() function
	 */
	a := new([5]int)
	b := new([30]int)[0:10]
	c := new([100]int)[0:30]

		// Appending the values to the slice
	c = append(b, []int{10, 20, 30, 40, 50}...)
	

	// Printing the type, length, capacity and value
	fmt.Printf("a - Type: %T, Length: %d, Capacity: %d, Value: %v\n",
		a, len(a), cap(a), a)

	fmt.Printf("b - Type: %T, Length: %d, Capacity: %d, Value: %v\n",
		b, len(b), cap(b), b)

	fmt.Printf("c - Type: %T, Length: %d, Capacity: %d, Value: %v\n",
			c, len(c), cap(c), c)
}