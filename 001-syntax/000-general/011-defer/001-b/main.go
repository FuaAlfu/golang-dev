package main

import "fmt"

//func (r receiver) identifier(parameters) (return(s)) { code }
func main() {
	defer foo()
	bar()
}

func foo() { fmt.Println("FOO") }

func bar() { fmt.Println("BAR") }
