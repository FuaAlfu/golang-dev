package main

import (
	"fmt"
	"time"
)

func sayHi(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println("hi there! ", i, source)
	}
	time.Sleep(time.Millisecond * 1)
}

func main() {
	go sayHi("goroutine")
	sayHi("main function")
}
