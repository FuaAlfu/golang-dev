package main

import "fmt"

func main() {
	/*
		memory allocated of this array
		is 48 byte
	*/
	var languages [3]string
	languages[0] = "Russian"  //16byte
	languages[1] = "French"   //16 byte
	languages[2] = "Japanese" //16 byte

	//iterate over the array of strings
	for _, language := range languages {
		fmt.Println(language)
	}
	println("---")
	//-----------------------------------

	//Declare an array of 4 integers that initilized with some values..
	numbers := [5]int{0, 10, 20, 30, 40}

	//iterate over the array of numbers.
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
	println("---")
	//----------------------------------

	//not good practice :: do not use it
	//look into the arrays elements to set up the size
	numbers2 := [...]int{0, 10, 20, 30, 40}
	for i := 0; i < len(numbers2); i++ {
		fmt.Println(i, numbers2[i])
	}
	println("---")

	//set up elements by this syntax
	numbers3 := [5]int{2: 0, 1: 10, 0: 20, 4: 30, 3: 40}
	for i := 0; i < len(numbers3); i++ {
		fmt.Println(i, numbers3[i])
	}
}
