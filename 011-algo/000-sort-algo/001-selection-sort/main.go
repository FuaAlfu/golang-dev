package main

import "fmt"

func selectionSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[length-i-1], data[maxIndex] = data[maxIndex], data[length-i-1]
	}
	for _, v := range data {
        fmt.Println(v)
    }
}

func main(){
	s := []int{7, 64, 9, 2, 1, 888, 761, 90}
	selectionSort(s)
}