package main

import (
	"fmt"

)

func recursivefunction() {
	fmt.Println("re..")
	recursivefunction()
}

func main() {
	recursivefunction()
}
