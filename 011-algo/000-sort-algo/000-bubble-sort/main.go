package main

import "fmt"

func bubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
				//fmt.Println(data) //for testing
			}
		}
	}
    for _, v := range data {
        fmt.Println(v)
    }
}

func main() {
	s := []int{7, 64, 9, 2, 1, 888, 761, 90}

	bubbleSort(s)
	fmt.Println("bubbleSort")
}
