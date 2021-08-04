package main

import (
	"fmt"

)

var count = 0

func recursive(limit int) {
	count++
	if count <= limit {
		fmt.Println(count)
		recursive(limit)
	}
}

func main() {
	l := 30
	recursive(l)
}
