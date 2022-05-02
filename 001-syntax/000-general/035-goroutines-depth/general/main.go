package main

import(
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

var counter int

func work() {
    time.Sleep(time.Second)

    for i := 0; i < 1e10; i++ {
        counter++
		fmt.Printf("count: %d", counter)
		if counter > 100{
			break
		}
    }

    wg.Done()
}

func main() {
	for i := 0; i < 20; i++ {
        wg.Add(1) // increases WaitGroup
        go work() // calls a function as goroutine
    }

    wg.Wait() // waits until WaitGroup is <= 0
}