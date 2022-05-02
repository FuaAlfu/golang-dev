package main

import(
	"fmt"
)

func Add(a,b int)int{
	return a + b
}

func Substruct(a,b int)int{
	return a - b
}

func main() {
	fmt.Println(Add(1,5))
	fmt.Println(Substruct(1,5))
}