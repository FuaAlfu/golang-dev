package main

import (
	"fmt"
	"math/rand"
	"time"
)

var try int
var guess int

func guessNumber(randomNumber int) {
	try += 1
	fmt.Println("try #: ", try)
	fmt.Println("Guess a number 0 - 10")
	fmt.Scanf("%d", &guess)
	if guess == randomNumber {
		fmt.Printf("successful, took you %d trys", try)
	} else {
		fmt.Println("try again")
		guessNumber(randomNumber)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	randomNumber := rand.Intn(11)
	fmt.Println("random generated number", randomNumber)
	guessNumber(randomNumber)
}
