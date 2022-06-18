package main

import "fmt"

func countingSort(data []int, maxValue int){
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)

	sortedIndex := 0
	length := len(data)

	for i := 0; i < length; i++ {
		bucket[data[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			data[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	for _, v := range data {
        fmt.Printf("%v ", v) 
    }
}

func main() {
	s := []int{7, 64, 9, 2, 1, 888, 761, 90}
	countingSort(s,888)
}