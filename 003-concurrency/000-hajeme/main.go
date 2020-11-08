package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("os\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	//here to run concurrency by.. Using go keyword
	go boo()
	hoo()

}

func boo() {
	for i := 0; i < 10; i++ {
		fmt.Println("boo: ", i)
	}
}

func hoo() {
	for i := 0; i < 10; i++ {
		fmt.Println("hoo: ", i)
	}
}
