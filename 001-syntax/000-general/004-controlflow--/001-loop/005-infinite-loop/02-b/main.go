package main

import (
	"fmt"
	"time"
)

func main() {
	loopTimer := time.NewTimer(time.Second * 9)

	for {
		fmt.Println("inside the infinite loop!")

		<-loopTimer.C
		break
	}
}
