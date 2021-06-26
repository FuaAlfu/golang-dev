package main

import "fmt"

func main() {
	value := "cow;dog;cat;goat"

	//substring from index 4 to length of string.
	substring1 := value[4:len(value)]

	//substring from index 8 to length of string.
	substring2 := value[8:len(value)]

	fmt.Println(substring1)
	fmt.Println(substring2)
	fmt.Println(value)
}
