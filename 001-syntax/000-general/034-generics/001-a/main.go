package main

import "fmt"

// type Number interface{
// 	int64 | float64
// }

type Number interface{
	int64 
}

func sumInt(n ...int64)int64{
	var s int64

	//range over ..
	for i,v := range n{
		s += v
		fmt.Println(i,v)
	}
	return s
}

func sumNumber[N Number](sn ...N) N{
	var s N
	for _,v := range sn{
		s += v
	}
	return s
}

func main(){
	sInt := []int64{5,99,641,87,9,42}

	fmt.Printf("non-generic: %v \n", sumInt(sInt...))
	fmt.Printf("generic: %v \n", sumNumber(sInt...))
}