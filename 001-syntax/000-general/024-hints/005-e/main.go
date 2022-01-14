package main

import "fmt"

var people []string

func rangeOver(x []string){
	for i := 0; i < len(x); i++{
		fmt.Println(x[i])
	}
}


func main() {
	people = []string{"john","ben","red-one"}
	fmt.Println("[", people[1], "] is going to buy the lunch for today")
	people = append(people,"doe")
	rangeOver(people)
	println("---")
	fmt.Println("people:\tvalue of [", people, "]\taddress of[", &people, "]")
}