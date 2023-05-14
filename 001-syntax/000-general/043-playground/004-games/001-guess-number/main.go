package main

import (
	"fmt"
	"math/rand"
	"time"

)

func main() {
	generateNumber()
}

func generateNumber(){
	// Generate a random number
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	fmt.Println("Welcome to the Guess the Number game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Can you guess it?")

	// Keep track of the number of guesses
	numGuesses := 0

	for {
		var guess int
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		numGuesses++

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", numGuesses)
			break
		}
	}
}