package main

import "fmt"

type (
	node struct {
		data int //use a single number
		next *node
	}

	linkList struct {
		head   *node
		length int //the long of the list
	}
)

//defined a reciver method
func (l *linkList) prepend(n *node) {
	second := l.head     //put head in a temp holder
	l.head = n           //set the new node
	l.head.next = second //let the new head point to the old head
	l.length++           //increase the length
}

func main() {
	mylist := linkList{}
	node1 := &node{data: 23}
	node2 := &node{data: 11}
	node3 := &node{data: 8}
	node4 := &node{data: 97}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)

	fmt.Println(mylist)
}
