package main

import (
	"fmt"
	"time"

)

/*
Would receive from ping channel
Wait for 1 second
Then send to pong channel
*/
func pinger(ping <-chan string, pong chan<- string) {
	for m := range ping {
		printAndDelay(m)
		pong <- "pong"
	}
}

/*
Would receive from pong channel
Wait for 1 second
Then send to ping channel
*/
func ponger(ping chan<- string, pong <-chan string) {
	for m := range pong {
		printAndDelay(m)
		ping <- "ping"
	}
}

func printAndDelay(msg string) {
	fmt.Println(msg)
	time.Sleep(time.Second)
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	// Player A
	go pinger(ping, pong)
	// Player B
	go ponger(ping, pong)

	// Player A starts the game
	ping <- "ping"

	for {
	}
}
