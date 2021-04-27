package main

import "fmt"

type (
	person struct {
		first string
		last  string
	}
	secretAgent struct {
		person
		ltk bool
	}
)

//Method , Type secretAgent .. Attach method to type Value
func (s secretAgent) speak() {
	fmt.Println("I'm", s.first, s.last, " - the secretAgent speak")
}
func (p person) speak() {
	fmt.Println("I'm", p.first, p.last, " - the person speak")
}

//Interface
type human interface {
	speak()
}

func bar(h human) {
	//Using switch to check out the types
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrr", h.(person).first) //assert type person
	case secretAgent:
		fmt.Println("I was into barrr", h.(secretAgent).first) //assert type secretAgent
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

//func (r receiver) identifier(parameters) (return(s)) { code }
func main() {

	sac1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sac2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "MoneyPenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr",
		last:  "Master",
	}

	fmt.Println(sac1)
	sac1.speak()
	sac2.speak()

	fmt.Println(p1)

	bar(sac1)
	bar(sac2)
	bar(p1)

	//conversions
	var x hotdog = 32
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int  //type hotdog = int
	y = int(x) //this is conversion,.. conversion type to..
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
