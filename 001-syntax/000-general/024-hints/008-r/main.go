package main

import(
	"fmt"
	// "log"
)

func message(a int){
	for i := 0; i < a; i++{
	if i % 2 == 0{
		fmt.Println("fizz")
	}else if i % 3 == 1{
		fmt.Println("buzz")
	}
	}
}

func anotherMessage(a int){
	for i := 0; i < a; i++{
		vO := i % 2 == 0
		vT := i % 3 == 1
	switch {
	case (vO):
		fmt.Println("fizz")
	case (vT):
		fmt.Println("buzz")
	default:
		fmt.Println("none")
		}
	}
}

func main() {
	// if err != nil{
	// 	log.Fatal(err)
	// }
	message(30)
	println("---")
	anotherMessage(10)
}