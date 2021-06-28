package main

import "fmt"

func main() {
	value := "GOLang"
	//Omit start index, this is the same as zero.
	substring1 := value[:2]
	substring2 := value[:3]

	//Test the first substring.
	if substring1 == "GO" {
		fmt.Println(true)
	} else if substring2 == "Lan" {
		//Test the second substring. false
		fmt.Println(true)
	}
}
