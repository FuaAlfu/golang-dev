package main

import "fmt"

func main() {
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

	//Comma OK idiom..   checkout a stored value [v,ok]
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

	println("===!===!===")
}
