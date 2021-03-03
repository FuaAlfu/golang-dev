package main

import "fmt"

func main() {
	//creating and constructing an anonymous struct
	randomCar 		maker   string
		model   string
		color   string
		mileage int
	}{
		maker:   "mercedes-benz",
		model:   "G63",
		color:   "red",
		mileage: 0,
	}

	fmt.Println("my car's color is", randomCar.color)
}
