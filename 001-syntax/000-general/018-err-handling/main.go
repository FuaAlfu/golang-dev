package main

import "fmt"

func main() {
	/*
		type error interface{  == struct
		Error() string
		}

		any method return a string ..
	*/
	n, err := fmt.Println("Hello")
	//common fashon / throw
	//here we checking for the errors
	if err != nil {
		fmt.Println(err)
	}

	//6 == .. h e l l o
	fmt.Println(n)
}
