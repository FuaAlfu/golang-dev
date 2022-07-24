package main

import "fmt"

func strike(s []int){
	for i := 0; i < len(s); i++{
		//sum := s[i+1]
		sum := append(s,i)
		fmt.Println(sum)
		fmt.Printf("sum %d ",sum)
	}
}

func main(){
	s := []int{22,9,7,321,90,71,70,62}
	strike(s)
	//fmt.Println(strike(s))
}