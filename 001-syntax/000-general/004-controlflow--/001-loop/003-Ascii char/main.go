package main

import "fmt"

func main() {
		//print ascii char..
		for i := 33; i <= 122; i++{
			fmt.Printf("%v\t%#U\n",i,i)
	}
}