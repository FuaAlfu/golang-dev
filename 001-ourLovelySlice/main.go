/*
.

slice
*/
package main

import "fmt"

func main() {
	//array
	var money [3]string
	money[0] = "gold"
	money[1] = "silver"
	money[2] = "bronze"

	//slice
	furit := []string{"apple", "orange"}
	println(furit)
	fmt.Println("===!===!===")

	//make() a slice
	furits := make([]string, 5)
	furits[0] = "apple"
	furits[1] = "banana"
	furits[2] = "orange"
	furits[3] = "beach"
	furits[4] = "strwberry"

	inspectSlice(furits)
	fmt.Println("===!===!===")

	//zero value
	var list []int
	fmt.Println(list, " this is the zero value of a slice")
	fmt.Printf("%T", list)
	fmt.Println()

	//empty slice
	s := []int{}
	fmt.Println(s, " this an empty slice")
	fmt.Printf("%T", s)

}

func inspectSlice(slice []string) {
	fmt.Printf("length[%d] Capacity[%d]\n ", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}
