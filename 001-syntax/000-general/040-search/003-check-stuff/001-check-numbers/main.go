package main

import(
	"fmt"
)

func checkNumbers(n int){
	for i := 0; i < n; i++{
		if i % 2 == 0{
			fmt.Printf("whe have even number: %d \n",i)
		}else{
			fmt.Printf("whe have odd number: %d \n",i)
		}
	}
}

func main(){
	myNumber := 22
	checkNumbers(myNumber)
}