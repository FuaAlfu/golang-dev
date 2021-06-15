package main

import "fmt"

//Queue
type Queue struct {
	items []int
}

//enqueue :: add a value at the back [end]
func (q *Queue) enqueue(i int) {
	q.items = append(q.items, i)
}

//dequeue :: remove a value at the back [end], and returns the removed value
func (q *Queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.enqueue(5)
	myQueue.enqueue(15)
	myQueue.enqueue(25)
	fmt.Println(myQueue)
	myQueue.dequeue()
	fmt.Println(myQueue)
}
