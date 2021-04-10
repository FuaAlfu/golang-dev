package main

import "fmt"

func main() {

	/*
	   send:    chan<-
	   receive: <-chan

	   closing channel by looping
	*/
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c) //go bar(c)

	fmt.Println("about to exit")
}

//send
//these chan became a new memory area
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

//receive
func bar(c <-chan int) {
	//print the value
	//without go keyword it will stuck
	for v := range c {
		fmt.Println(v)
	}
	//fmt.Println(<-c)
}
