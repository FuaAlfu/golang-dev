package main

import "fmt"

func firstIndex(a[]int){
	first := a[0]  //a[:0]
	fmt.Printf("first index of array: %d\n of address %p\n", first, &first)
}

func lastIndex(a ...int){
	last := a[len(a)-1]
	fmt.Printf("last index of array: %d \n of address %p\n", last, &last)
}

func removeLastElement(a[]int){
	newArray := a[:len(a)-1]
	fmt.Printf("currnt array: %d \n of address %p\n", newArray, &newArray)
}

func removeFirstElement(a[]int){
	newArray := a[1:]
	fmt.Printf("currnt array: %d \n of address %p", newArray, &newArray)
}

func main(){
	myArray := []int{8,912,23,7}
	firstIndex(myArray)
	lastIndex(myArray...)
	removeLastElement(myArray)
	removeFirstElement(myArray)
}