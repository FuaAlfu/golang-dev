package main

import "fmt"

func main() {
	value := "wake;sleep"

	//specify just the first index.
	substring := value[5:]
	fmt.Println(substring)
}
