package main

import "fmt"

func increment(x int) func() (int, bool) {
	i := -1
	return func() (int, bool) {
		if i >= x {
			return 0, true
		}
		i++
		return i, false
	}
}

func revers(x int) func() (int, bool) {
	i := 0
	return func() (int, bool) {
		if i >= x {
			return 0, true
		}
		i--
		return i, false
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for j := 9; j < 10; j-- {
			if j == 0 {
				break
			}
			fmt.Println(j)
		}
	}
	//-----

	myFunc := increment(5)
	for i, eof := myFunc(); !eof; i, eof = myFunc() {
		fmt.Println(i)
	}
	myAnotherFunc := revers(10)
	for i, eof := myAnotherFunc(); !eof; i, eof = myAnotherFunc() {
		fmt.Println(i)
	}
	println("---")
}
