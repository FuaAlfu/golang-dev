package main

import "fmt"

type basacon struct {
	first   string
	another string
	price   float64
}

type marketing struct {
	basacon
	readyToSell bool
}

func main() {
	//initialize the value of type {struct} inside another...
	m := marketing{
		basacon: basacon{
			first:   "Sony",
			another: "Bravia",
			price:   23.888,
		},
		readyToSell: true,
	}
	fmt.Println(m)

	//another way to access the value
	fmt.Println(m.first, m.another, m.price, m.readyToSell)
}
