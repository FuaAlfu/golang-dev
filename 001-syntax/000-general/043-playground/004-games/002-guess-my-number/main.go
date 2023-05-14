package main

import (
	"fmt"
	"math/rand"
	"time"

)

func main() {
	guessMyNumberMyLovelyComputer()
}

func guessMyNumberMyLovelyComputer(){
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("Think of a number between 1 and 100.")

	min := 1
	max := 100
	numGuesses := 0
	found := false

	for !found {
		guess := rand.Intn(max-min+1) + min
		fmt.Printf("Is your number %d? (y/n): ", guess)

		var response string
		fmt.Scanln(&response)

		switch response {
		case "y":
			fmt.Printf("Yay! I guessed your number (%d) in %d attempts!\n", guess, numGuesses+1)
			found = true
		case "n":
			fmt.Println("Is it higher or lower?")
			fmt.Print("Enter 'h' for higher or 'l' for lower: ")

			var hint string
			fmt.Scanln(&hint)

			switch hint {
			case "h":
				min = guess + 1
			case "l":
				max = guess - 1
			default:
				fmt.Println("Invalid input. Please try again.")
			}

			numGuesses++
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}