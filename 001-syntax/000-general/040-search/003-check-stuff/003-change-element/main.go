package main

import(
	"fmt"
)
func changeElement(s []int, e int, newElement int){
	for i := 0; i < len(s); i++{
		    fmt.Printf("we had this element: %d \n",s[e])
			s[e] = newElement
			fmt.Printf("now we got: %d", newElement)
			break;
	}
}

func main(){
	s := []int{33,88,9,6,32,123,321}
	changeElement(s,5,99)
}