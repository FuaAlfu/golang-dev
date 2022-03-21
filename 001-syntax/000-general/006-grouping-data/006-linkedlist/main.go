package main

import "fmt"

type(
	node struct {
		data int
		next *node
	}

	linkedList struct {
		length int
		head   *node
		tail   *node
	}
)

func (l linkedList) Len() int {
    return l.length
}

func (l linkedList) Display() {
    for l.head != nil {
        fmt.Printf("%v -> ", l.head.data)
        l.head = l.head.next
    }
    fmt.Println()
}

func (l *linkedList) PushBack(n *node) {
    if l.head == nil {
        l.head = n
        l.tail = n
        l.length++
    } else {
        l.tail.next = n
        l.tail = n
        l.length++
    }
}

func main() {
	mylist := linkedList{}
	n1 := &node{data: 5}
	n2 := &node{data: 10}
	n3 := &node{data: 15}
	n4 := &node{data: 20}

	mylist.PushBack(n1)
	mylist.PushBack(n2)
	mylist.PushBack(n3)
	mylist.PushBack(n4)

	mylist.Display()

	fmt.Println(mylist)
}