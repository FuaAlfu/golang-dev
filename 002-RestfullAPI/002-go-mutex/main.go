package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	cridite int
)

func main() {
	/*
		banking app
	*/
	fmt.Println("banking around\n")

	cridite = 500.00

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(670, &wg)
	go desposit(2600, &wg)
	wg.Wait()

	fmt.Printf("your new balance %d\n", cridite)
}

func desposit(v int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("desposit %d to cridite %d\n", v, cridite)
	cridite += v
	mutex.Unlock()
	wg.Done() //goroutine is done
}

func withdraw(v int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("withdraw %d from account with balance  %d\n", v, cridite)
	cridite -= v
	mutex.Unlock()
	wg.Done() //goroutine is done
}
