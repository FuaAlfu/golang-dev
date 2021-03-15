package main

import "fmt"

/*this struct cost [8] byte
we could reduce struct cost to ...
(if we want to)
*/

type example struct {
	counter int
	flag    bool
	pi      float32
}

func main() {
	//zero value
	var e1 example

	//another way to type a zero value
	e2 := example{} //, or e2 := return example{}

	//hase value: a struct literal
	e3 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
	println("---")
	fmt.Println(e3.pi)
}
