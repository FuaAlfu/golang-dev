package main

import "fmt"

type Queue struct {
    items []int
}

func (q *Queue) Enqueue(data int) {
    q.items = append(q.items, data)
}

func (q *Queue) Dequeue() {
    q.items = q.items[1:]
}

func (q *Queue) Front() (int, error) {
    if len(q.items) == 0 {
        return 0, fmt.Errorf("queue is empty")
    }

    return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
    if len(q.items) == 0 {
        return true
    } else {
        return false
    }
}

func (q *Queue) Print() {
    for _, item := range q.items {
        fmt.Print(item, " ")
    }
    fmt.Println()
}

func main() {
    myQueue := Queue{[]int{}}
    myQueue.Print()
    fmt.Println(myQueue.IsEmpty())

    myQueue.Enqueue(100)
    myQueue.Enqueue(2000)
    myQueue.Enqueue(30000)

    myQueue.Print()
    myQueue.Dequeue()
    myQueue.Print()
}