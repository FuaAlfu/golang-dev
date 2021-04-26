package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
		age  int
	}

	BankInfo struct {
		Person
		credit float64
	}
)

func Println2(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Print("a boolean value: ", x.(bool))
	case int:
		fmt.Print("an integer value: ", x.(int))
	case float64:
		fmt.Print(x.(float64))
	case complex128:
		fmt.Print(x.(complex128))
	case string:
		fmt.Print(x.(string))
	case Person:
		fmt.Print(x.(Person))
	case chan int:
		fmt.Print(x.(chan int))
	default:
		fmt.Print("Unknown type")
	}

	fmt.Print("\n")
}

func Println3(x interface{}) {
	fmt.Printf("type is '%T', value: %v\n", x, x)
}

func main() {

	y := 15
	z := true
	p1 := BankInfo{
		credit: 33245673.00,
		Person: Person{
			name: "john",
			age:  30,
		},
	}

	Println2(y)
	Println3(y)
	println("---")
	Println2(z)
	Println3(z)
	println("---")
	Println2(p1)
	Println3(p1)
}
