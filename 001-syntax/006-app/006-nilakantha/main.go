package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	ANSIClearScreenSequence       = "\033[H\033[2J"
	ANSIFirstSlotScreenSequence   = "\033[2;0H"
	ANSISecondSlotScreeenSequence = "\033[3;0H"
)

var pichan chan float64 = make(chan float64)

var computationDon chan bool = make(chan bool, 1)

var termsCount int

func printCalculation() {
	fmt.Println(ANSIClearScreenSequence)
	fmt.Println(ANSIFirstSlotScreenSequence, "computed value of pi:\t\t", <-pichan)
	fmt.Println(ANSISecondSlotScreeenSequence, "# of Nilakantha term:\t\t", termsCount)
}

func pi(n int) float64 {
	ch := make(chan float64)
	for k := 1; k <= n; k++ {
		go nilakanthaTerm(ch, float64(k))
	}
	f := 3.0
	for k := 1; k <= n; k++ {
		termsCount++
		f += <-ch
		pichan <- f
	}

	computationDon <- true
	return f
}

func nilakanthaTerm(ch chan float64, k float64) {
	j := 2 * k
	if int64(k)%2 == 1 {
		ch <- 4.0 / (j * (j + 1) * (j + 2))
	} else {
		ch <- -4.0 / (j * (j + 1) * (j + 2))
	}
}

func main() {
	/*
		the more neila contract terms we calculate the more
		accurate our computed value of pi will be..
	*/

	ticker := time.NewTicker(time.Microsecond * 108)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go pi(5000)

	go func() {
		for range ticker.C {
			printCalculation()
		}
	}()

	for {
		select {
		case <-computationDon:
			ticker.Stop()
			fmt.Println("program done calculating pi..")
			os.Exit()

		case <-interrupt:
			ticker.Stop()
			fmt.Println("program interrupt by the user,..")
			return
		}
	}
}
