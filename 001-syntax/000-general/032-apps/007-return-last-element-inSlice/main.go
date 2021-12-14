package main

import "fmt"

func returnLastElement(a ...int) {
	fmt.Println(a[len(a)-1:][0])
}

func returnLastElementS(a []int) {
	fmt.Println(a[len(a)-1])
}

func main() {
	s := []int{3, 88, 51, 999, 1, 12}
	returnLastElement(s...)
	returnLastElementS(s)
}
