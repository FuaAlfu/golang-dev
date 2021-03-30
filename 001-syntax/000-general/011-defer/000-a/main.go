package main

import "fmt"

func runing() {
	fmt.Println("runing..")
}

func main() {
	defer runing()
	fmt.Println("hello")
}
