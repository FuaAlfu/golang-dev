package main

import (
	"fmt"
	"math"

)

func getMaxInt32() int {
	return math.MaxInt32
}

func main() {
	fmt.Printf("Type of math.MaxInt32 is %T\n", math.MaxInt32)
	fmt.Println("Value of math.MaxInt32:", math.MaxInt32)
	fmt.Println("Value of math.MaxInt32:", getMaxInt32())
}
