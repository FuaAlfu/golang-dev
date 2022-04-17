package main

import(
	"fmt"
)

func sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, sum(tail))
		}
	}

	return sums
}

//generic
func reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return reduce(numbers, add, 0)
}


func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return reduce(numbers, sumTail, []int{})
}

func main() {
	n1 := []int{11,22,33,44,55}
	fmt.Println(sum(n1))
	fmt.Println(sumAllTails(n1))
	println("---")
	fmt.Println(Sum(n1))
	fmt.Println(SumAllTails(n1))

}