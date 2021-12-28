package main

import "fmt"

type Slice struct {
	s [][]int
}

func main() {
	z := Slice{
		s: [2][3]int{
			{33, 66, 0},
			{52, 98, 5},
		},
	}
	//---test
	slice := [2][3]int{
		{33, 66, 0},
		{52, 98, 5},
	}
	fmt.Println(z.s)
	println("===")
	fmt.Println(slice)
}
