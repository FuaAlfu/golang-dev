package main

import (
	"fmt"

)

type Queue interface {
	Enqueue(value interface{})
	Dequeue() interface{}
	IsEmpty() bool
	Size() int
}

type queue struct {
	values []interface{}
}

func (q *queue) Enqueue(value interface{}) {
	q.values = append(q.values, value)
}

func (q *queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	value := q.values[0]
	q.values = q.values[1:]

	return value
}

func (q *queue) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *queue) Size() int {
	return len(q.values)
}

func main() {
	myQueue := &queue{}

	myQueue.Enqueue(10)
	myQueue.Enqueue("Hello")
	myQueue.Enqueue(3.14)

	fmt.Println("Queue size:", myQueue.Size())

	value := myQueue.Dequeue()
	fmt.Println("Dequeued value:", value)

	fmt.Println("Is queue empty?", myQueue.IsEmpty())
}
