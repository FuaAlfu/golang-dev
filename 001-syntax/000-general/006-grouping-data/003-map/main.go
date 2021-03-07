package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		including sorting map
	*/

	//Unordered List
	//[Key]Value{}
	m := map[string]int{
		//Key:Value,
		"Fua":    30,
		"Miss N": 27,
	}
	fmt.Println(m)

	//print up the Value..
	fmt.Println(m["Fua"])
	fmt.Println(m["Miss N"])

	//Value not stored in map
	fmt.Println(m["Huja"])

	//Comma OK idiom..   check stored value [v,ok]
	v, ok := m["Huja"]
	fmt.Println(v)
	fmt.Println(ok)

	//a Common way to type Comma ok idiom in Go..
	if v, ok := m["Miss N"]; ok {
		fmt.Println("This is the if print", v)

		//OutPut will be false..
		//try it with another les..
		if v, ok := m["Gujaa Gujaa"]; ok {
			fmt.Println("This is the if print", v)
		}
	}

	//add new element
	m["platinum"] = 300

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Printf("%v\n", v)
	}
	println("---")

	//sorting map
	printSortedMap(m)
}

func printSortedMap(m map[string]int) {

	keys := make([]string, len(m))

	for key, _ := range m {

		keys = append(keys, key)
	}
	sort.Strings(keys) //here

	for _, v := range keys {
		continue
		fmt.Println("the..", v, "is", m[v])
	}
}
