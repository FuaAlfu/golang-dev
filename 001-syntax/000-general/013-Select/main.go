package main

import "fmt"

func main() {
	/*
	send:    chan<-
	receive: <-chan

	the main pieces here\
	channels is block
	closing channel by looping (Range)

	receive pull values from the channel
	select statement
	*/

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Println("about to exit")

}

func receive(e, o, q <-chan int) {
	//infit loop for keyword only
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel ", v)
		case v := <-o:
			fmt.Println("from the odd channel ", v)
		case v := <-q:
			fmt.Println("from the quit channel ", v)
			//	close(q)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
