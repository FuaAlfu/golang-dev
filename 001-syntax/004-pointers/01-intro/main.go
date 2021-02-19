package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Zero"
	//or we could use “dereferencing“ to get the value
	// (*p).name = "Zero"
}

func main() {
	/*pointer semantic -> we got and keep one copy of data
	and data allowed to share by everybody
	*/

	//default Anonymous func
	f := func() int {
		return 0
	}()

	//using ampersand mark to print out the address..
	fmt.Println(f)
	fmt.Println(&f)
	print("----------------\n")

	//constractor
	p1 := person{
		name: "Div",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}
