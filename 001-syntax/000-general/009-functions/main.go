package main

import "fmt"

//func (r receiver) identifier(parameters) (return(s)) { ... }
func main() {
	hi()
	hellow("Smoky")
	meshMush()
	fmt.Println("====================")
	moshiMoshi()
	s1 := "Natolah"
	fmt.Println(s1)
	fmt.Println("---")
	x, y := mouse("Ian", " Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func hi() {
	fmt.Println("Hello and welcome.. ")
}

//everything in go is pass by value, what you see is what you get .(we see we ge).
func hellow(s string) {
	fmt.Println("Mr. ", s)
}

//arg and using a return method, impliment  a method function `Sprint`.
func woo(st string) string {
	return fmt.Sprint("Helloo woo", st)
}

//Two types..a & b .. multiple returns
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `,says: "Hello"`)
	b := false
	return a, b
}

//searching log by mapping data
func meshMush() int {
	m := map[string]int{
		"storm": 1,
		"cloud": 2,
		"rain":  3,
		"wind":  4,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	return 0
}

//searching log by looking inside element within data
func moshiMoshi() {
	badAsses := []string{
		"go",
		"ho",
		"bo",
	}
	//	var chick = "no"
	for i, val := range badAsses {
		fmt.Println(i, val)
		//if chick == x {
		//	fmt.Println(x)
		//	break
		//}

	}
}
