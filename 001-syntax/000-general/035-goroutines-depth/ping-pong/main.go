package main
/*
for testing
*/
import (
	"fmt"
	"time"

)

var count int64

func loop(msg string) {
	for {
		count += 1
		fmt.Println(msg)
		time.Sleep(time.Second)
		if count == 16 {
			fmt.Printf("finished: %d time ",count)
			break
		}
	}
}

func main() {
	go loop("First")
	loop("Second")
}
