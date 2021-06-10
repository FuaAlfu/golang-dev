package main

import "fmt"

type (
	Book struct {
		title string
		pages int
		price float64
	}
	book []Book
)

func main() {
	b := book{
		Book{
			title: "lord of the love",
			pages: 180,
			price: 23.00,
		},
	}

	fmt.Println(b)
}
