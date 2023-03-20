package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel to receive results from the goroutines
	resultChan := make(chan int)

	// launch two goroutines that generate random numbers
	go generateRandomNumbers(resultChan)
	go generateRandomNumbers(resultChan)

	// wait for the goroutines to finish and collect the results
	result1 := <-resultChan
	result2 := <-resultChan

	// print the results
	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
}

func generateRandomNumbers(resultChan chan int) {
	// generate a random number and send it to the channel
	result := rand.Intn(100)
	resultChan <- result

	// simulate some work
	time.Sleep(1 * time.Second)
}
