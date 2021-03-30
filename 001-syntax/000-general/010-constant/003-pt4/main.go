package main

import "fmt"

type Duration int64
type handle int

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisocond          = 1000 * Microsecond
	Second               = 1000 * Millisocond
	Munite               = 60 * Second
	Hour                 = 60 * Munite
)

const (
	one = 1000
	two = one * 1000
)

func hoo(h handle) {
	fmt.Println(h)
}

func time(time int) {
	fmt.Println(time)
}

func main() {
	fmt.Println("const")
	var h handle //zero value
	hoo(h)
	println("---")
	fmt.Println("time: ", Nanosecond, Microsecond, Millisocond, Second, Munite, Hour)
	println("---")
	time(two)
}
