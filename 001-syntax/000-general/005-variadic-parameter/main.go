package main

import "fmt"

//func (r receiver) identifier(parameters) (return(s)) { code }
func main() {
	y := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("the total is", y)
	println("----")

	//add lexical element <...> after the slice
	z := []int{22, 77, 31}
	fmt.Println(sum(z...))
}

//Variadic parameters (x ...int) three dots  … "lexical element" Type slice of = []int
//lexical symbol: ... = slice
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position:,", i, " we are now adding ", v, "to the total which is now")
	}
	fmt.Println("the total is: ", sum)
	return sum
}
