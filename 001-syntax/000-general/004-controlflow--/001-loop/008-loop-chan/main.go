package main

import (
	"fmt"
	"strconv"
)

func increment() chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- strconv.Itoa(i)
		}
		close(ch)
	}()
	return ch
}

func main() {
	for i := range increment() {
		fmt.Println(i)
	}
}
