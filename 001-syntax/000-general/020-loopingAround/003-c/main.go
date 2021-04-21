package main

import "fmt"

//var plus int  //Plus =0
func main() {

	/*var array [7] int
	fmt.Println(array)*/

	nums := []int{5, 6, 9}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Print("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index: ", i)
		}
	}
	v := map[string]string{"a": " apple", "b": " banana"}
	for c, s := range v {
		fmt.Printf("%s -> $s\n", c, s)
	}
	for j := range v {
		fmt.Println("key:", j)
	}
	for r, b := range "go" {
		fmt.Println(r, b)
	}
}
