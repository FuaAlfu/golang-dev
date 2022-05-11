package main

import (
	"fmt"
	"math/rand"
	"time"

)

func randomSlice(len int) []int {
	s := make([]int, len)
	for i := 0; i <= len-1; i++ {

		s[i] = rand.Intn(len)
	}
	return s
}

func run(len int){
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomSlice(len))
}

func main() {
	/*
	   random array of integers
	*/
	run(10)
}
