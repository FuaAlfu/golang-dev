package main

import "fmt"

type Slice struct{
	slice []int
}

func addElements(s []int,length int){
	for i := 0; i < length; i++{
		s = append(s,i)
		fmt.Printf("e: %d \n",s)
	}
}

func addInnerElements(s []int,length int){
	for i := 0; i < length; i++{
		s = append(s,i)
		fmt.Printf("e: %d \n",s)
		for j := 0; j < i; j++{
		s = append(s,j)
		fmt.Printf("e: %d \n",s)
		}
	}
}

func main(){
	s := Slice{
		slice : []int{},
	}

	addElements(s.slice,9)
	println("---")
	addInnerElements(s.slice,9)
}