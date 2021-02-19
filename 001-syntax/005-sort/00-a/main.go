package main

import (
	"fmt"
	"sort"
)

func main() {
	/**
	using sorts\
	sort.ints
	sort.Strings**/

	//unsorted slices
	i := []int{77, 88, 11, 22, 8, 56, 999, 1, 787}
	s := []string{"fua", "alfu", "Dr"}
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println("----------------------------")

	//sorting
	sort.Ints(i)
	fmt.Println(i)
	sort.Strings(s)
	fmt.Println(s)

	/*
	   OUTPUT:
	   [77 88 11 22 8 56 999 1 787]
	   [fua alfu Dr]
	   ----------------------------
	   [1 8 11 22 56 77 88 787 999]
	   [Dr alfu fua]

	   Process finished with exit code 0
	*/

}
