package main

import "fmt"

type person struct {
	name string
	age  int
}

var passAGE int

func passAge(a int) {
	passAGE = 21

	switch {
	case (a >= passAGE):
		fmt.Println("pass !")
	case (a < passAGE):
		fmt.Println("not pass !!")
	default:
		fmt.Println("default")

	}
}

func passAgeByBool(a int) {
	pass := a >= 21
	notPass := a < 21

	switch pass {
	case pass:
		fmt.Println("pass !")
	case notPass:
		fmt.Println("not pass !!")
	default:
		fmt.Println("default")

	}
}

func main() {
	p1 :=
		person{
			name: "juana",
			age:  22,
		}

	p2 :=
		person{
			name: "anna",
			age:  21,
		}

	passAge(p1.age)
	passAgeByBool(p2.age)

}
