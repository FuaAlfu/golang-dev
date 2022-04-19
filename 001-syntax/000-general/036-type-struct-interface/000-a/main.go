package main

import(
	"fmt"
)

type(
	Hole struct{
		person Person
		car Car
	}
	Person struct{
		name string
		age int
	}
	Car struct{
		brand string
		doors int
	}
	Colour interface{
		Clr() string
	}
)

func Clr(clr string) string{
	return clr
}

// func rangeOver(h Hole) {
// 	for i,v := range h{
// 		fmt.Println(i,v)
// 	}
// }

func main() {
	h := Hole{
		person: Person{
			name: "john doe",
			age: 22,
		},
		car: Car{
			brand: "x",
			doors: 2,
		},
	}

	//rangeOver(h)
	fmt.Println(Clr("red"))
	println(h.person.name)
}