package main

import "fmt"

//
type Stack struct {
	items []int
}

//push :: add a value at the back [end], it need to take in a stack as a pointer receiver
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

//pop :: remove a value at the back [end], and returns the removed value
func (s *Stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {

	/*
		stack
		toRemove := s.items[len(s.items)-1]
	*/
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.push(33)
	myStack.push(22)
	myStack.push(11)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)

}
