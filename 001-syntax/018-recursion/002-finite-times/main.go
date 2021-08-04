package main

import (
	"fmt"
)

var count = 0

func recursivemethod(limit int) {
	count++
	if count <= limit {
		fmt.Println(count)
		recursivemethod(limit)
	}
}

func main() {
	l := 30
	recursivemethod(l)
}
