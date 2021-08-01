package main

import "fmt"

type person struct {
	name string
}

func main() {

	//constructing
	p1 := person{
		name: "fua",
	}

	p2 := person{
		name: "mua",
	}

	fmt.Println(p1.name)
	fmt.Println(arrayReverse(p1.name))
	fmt.Println(p2.name)
	fmt.Println(arrayReverse(p2.name))

	//---------------------------------
	value1 := "cat"
	reversed1 := arrayReverse(value1)
	fmt.Println(value1)
	fmt.Println(reversed1)

	value2 := "abcde"
	reversed2 := arrayReverse(value2)
	fmt.Println(value2)
	fmt.Println(reversed2)
}

func arrayReverse(value string) string {
	/*
			 Convert string to rune slice.
		    This method works on the level of runes, not bytes.
	*/
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}
