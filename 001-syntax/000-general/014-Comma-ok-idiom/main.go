package main

import "fmt"

func main() {
	/*
	send:
	chan <-

	receive:
	<- chan

	comma ok idiom with select
	*/
	c := make(chan int)
	go func() {
		c <- 30
		//close
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	fmt.Println(c)
	fmt.Println(<-c)

	//double check..
	v, ok = <-c
	fmt.Println(v, ok)
}
