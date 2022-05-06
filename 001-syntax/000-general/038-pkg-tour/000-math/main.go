package main

import(
	"fmt"
	"math"
	"math/rand"
)

/*
this function will demonstrate the use of rand.Intn() function.
The rand.Intn() function returns a random number between 0-N.
N is the specified number in function rand.Intn()
*/
func readInt(i int){
	fmt.Println("your number: ",rand.Intn(i))
}

//find the square root
func ReadFloat(i float64){
	fmt.Printf("your number: %g",math.Sqrt(i))
}

func main() {
	number := 69
	number2 := 69.00
	pi := math.Pi
	n := math.Abs(number2)
	p := math.Pow(number2,number2)
	s := math.Sin(pi)

	fmt.Println(pi)
	fmt.Println(n)
	fmt.Println(p)
	fmt.Println(s)
	readInt(number)
	ReadFloat(number2)
}