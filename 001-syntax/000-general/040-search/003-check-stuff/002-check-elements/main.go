package main

import "fmt"

func checkElements(s []int){
	for i := 0; i < len(s); i++{
		if s[i] % 2 != 0{
			fmt.Printf("whe have odd element: %d \n",s[i])
		}else{
			fmt.Printf("whe have even element: %d \n",s[i])
		}
	}
}

func main() {
	s := []int{33,88,9,6,32,123,321}
	checkElements(s)
}