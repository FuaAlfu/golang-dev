package main

import "fmt"

func main() {
	s := []int{554, 7, 11}
	g := []int{6, 8, 222}

	//add new elements to a slice
	sa := append(s, 99, 202)
	sg := append(s, g...)
	fmt.Println(sa)
	fmt.Println(sg)

	//adding two large slices..
	appendAslice(sa, sg)
	println("=--_--")
	inspectSlice(s)
}

func appendAslice(s []int, s2 []int) {
	a := append(s, s2...)
	fmt.Println(a)
}

func inspectSlice(slice []int) {
	fmt.Printf("length[%d] Capacity[%d]\n ", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}
