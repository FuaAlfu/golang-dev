package main

import(
	
	"time"
	"math/rand"
)

func randomQuickSort(list []int, start, end int) {
	if end-start > 1 {
		mid := randomPartition(list, start, end)
		randomQuickSort(list, start, mid)
		randomQuickSort(list, mid+1, end)
	}
}

func randomPartition(list []int, begin, end int) int {
	rand.Seed(time.Now().UnixNano())
	i := begin + rand.Intn(end-begin)
	list[i], list[begin] = list[begin], list[i]
	return partition(list, begin, end)
}

func partition(list []int, begin, end int) (i int) {
	cValue := list[begin]
	i = begin
	for j := i + 1; j < end; j++ {
		if list[j] < cValue {
			i++
			list[j], list[i] = list[i], list[j]
		}
	}
	list[i], list[begin] = list[begin], list[i]
	return i
}

func main() {
	s := []int{7, 64, 9, 2, 1, 888, 761, 90}
	randomQuickSort(s,0,8)
}