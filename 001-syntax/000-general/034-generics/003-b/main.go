package main

import "fmt"

type comparable struct{
	f int64
}

func sunIntOrFloat[k comparable, v int64  | float64](m map[k]v)v{
	var s v
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	sInt := map[int64]int64{1:5,2:99,3:641,4:87,5:9,6:42}
	fmt.Printf("generic: %v \n", sunIntOrFloat(sInt))
}