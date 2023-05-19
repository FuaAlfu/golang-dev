package main

import (
	"fmt"

)

type Stack interface {
	Push(value interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
}

type stack struct {
	values []interface{}
}

func (s *stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	index := len(s.values) - 1
	value := s.values[index]
	s.values = s.values[:index]

	return value
}

func (s *stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *stack) Size() int {
	return len(s.values)
}

func main() {
	myStack := &stack{}

	myStack.Push(10)
	myStack.Push("Hello")
	myStack.Push(3.14)

	fmt.Println("Stack size:", myStack.Size())

	value := myStack.Pop()
	fmt.Println("Popped value:", value)

	fmt.Println("Is stack empty?", myStack.IsEmpty())
}
