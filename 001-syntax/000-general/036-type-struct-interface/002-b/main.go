package main

import(
	"fmt"
)

type(
	car interface {
		color() string
		speed() float64
	}

	unit struct{
		count []int64
		speed float64
		colour string
	}
)

func carUnit(c unit)*unit{
	return &c
}

// func colour(s string)string{
// 	return s
// }

func speed(f float64)float64{
	return f
}

func rangeOver(r ...int64)[]int64{
	for i, v := range r{
		return i,v
	}
	return r
}

func main() {
	u1 := unit{
		count: []int64{1,65,90,4},
		speed: 67.90,
		colour: "red",
	}

	fmt.Println(speed(u1.speed))
//	fmt.Println(color(u1.colour))
	fmt.Println(carUnit(u1))
	fmt.Println(rangeOver(u1.count...))
}