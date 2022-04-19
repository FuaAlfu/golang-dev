package main

import(
	"fmt"
)

type Data struct{
	key int
	val []int
}

func changeWithPointer(arr *[]int) {
	*arr = append(*arr, []int{4, 5, 6}...)
	fmt.Printf("changedWithPointer => %v\n", arr)
}

// func changeNoPointer(arr []int) {
// 	arr = append(arr, []int{7, 8, 9}...)
// 	fmt.Printf("changedNoPointer => %v\n", arr)
// }

func main() {
	data := Data{
		key: 1,
		val: []int{11,7,982},
	}

	//d := []int{11,7,982}

	changeWithPointer(&data.val)
	//changeNoPointer(data.val)
}