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
	for i, language := range languages {
		fmt.Println(i, language)
	}
}
