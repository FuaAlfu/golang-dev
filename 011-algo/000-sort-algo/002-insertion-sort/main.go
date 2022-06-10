package main

import "fmt"

func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			j := i - 1
			temp := data[i]
			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}
	}
	for _, v := range data {
        fmt.Printf("%v ", v) 
    }
}

func main(){
	s := []int{7, 64, 9, 2, 1, 888, 761, 90}
	InsertionSort(s)
}