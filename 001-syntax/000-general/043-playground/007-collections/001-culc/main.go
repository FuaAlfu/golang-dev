package main

import "fmt"

var taxRate float64

func culc(income float64) float64{
	return income * taxRate
}

func main() {
	taxRate = 0.40
	limit := 5000.00
	var inc float64
		fmt.Print("Enter your income: ")
		fmt.Scanln(&inc)
		if inc * taxRate >= limit{
			fmt.Printf("you were accepted: above than tax limit %v\n", limit)
		}else{
			fmt.Printf("you were not accepted: less than tax limit %v\n", limit)
		}
	fmt.Println("your tax is:",culc(inc))
	
}