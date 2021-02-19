package main

import "fmt"

func main() {
	m := map[string]int{
		"gold":   100,
		"silver": 75,
		"bronze": 35,
	}

	fmt.Println(m)
	println()

	//add new element
	m["platinum"] = 300

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Printf("%v\n", v)
	}

}
