package main

import (
	"fmt"
	"sort"

)

func printSortedMap(m map[string]string) {
	keys := make([]string, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, v := range keys {
		//continue
		fmt.Println("the.. ",v, ", is ", m[v])
	}
}

func main() {

	m := map[string]string{
		"a": "10",
		"b": "100",
		"c": "1000",
	}
	m["d"] = "10000"

	printSortedMap(m)
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// 	fmt.Printf("%v\n", v)
	// }
	// println("---")

	// for i := 0; i < len(m); i++{
	// 	fmt.Println(i,m[i][0])
	// }
}
