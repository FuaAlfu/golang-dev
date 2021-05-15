package main

import (
	"fmt"

)

func increment(inc *int) {
	*inc++
	fmt.Println("inc:\tvalue of [", inc, "] and address of [", &inc, "]\t point to [", *inc, "] .. ")
}

func main() {
	count := 999

	//our func increment the value[count] that the pointer point to..
	increment(&count)
	println("count:\tvalue of [", count, "] and address of [", &count, "]")
}
