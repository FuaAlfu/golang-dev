package main

import "fmt"

func main() {
	//Declare an array of 5 integers that is initialized by zero value
	var six [6]int
	var four1 [4]int

	//Declare an array of 5 integers that is initialized
	var five [5]int
	five[0] = 1
	five[1] = 2
	five[2] = 3
	five[3] = 4
	five[4] = 5

	//iterate over the array of integers
	for i, f := range five {
		fmt.Println(i, f)
	}
	println("---")
	//------------------------------------

	//Declare an array of 4 integers that is initialized
	four := [4]int{10, 20, 30, 40}
	//iterate over the array of integers
	for i, f := range four {
		fmt.Println(i, f)
	}
	println("---")

	//assign one array to the other
	//six = four  //mismach :: diffrent int type .. not work

	//assign another array to the other
	four1 = four
	fmt.Println(four1)
	fmt.Println(six)
	println("---")
	fmt.Println("printing five elements from array five:\t [", five, "]")
	fmt.Println("printing the address of fourth element from array five:\t [", &five[3], "]")
	fmt.Printf("\t\t The inner value of an array: %d\n", five)

}
