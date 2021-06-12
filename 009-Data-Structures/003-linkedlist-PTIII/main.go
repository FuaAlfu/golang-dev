package main

import "fmt"

type (
	node struct {
		data int //use a single number
		next *node
	}

	linkedList struct {
		head   *node
		length int //the long of the list
	}
)

//defined a pointer reciver method
func (l *linkedList) prepend(n *node) {
	second := l.head     //put head in a temp holder
	l.head = n           //set the new node
	l.head.next = second //let the new head point to the old head
	l.length++           //increase the length
}

//defined a value reciver method
func (l linkedList) printlistData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next //to update the loop
		l.length--
	}
	fmt.Printf("\n")
}

//delete a node with a given value
func (l *linkedList) deleteWithValue(value int) {
	//handle err cases at run time
	//case 1
	if l.length == 0 {
		return
	}

	//case 2 :if the enterd value was the heaser ..we can handle it
	if l.length.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	//scan over the nodes
	//compare the data of the next node..
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		//case 3 :if input values does not exist inside the linkedlist
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next //update the loop :: looping till find the match
	}
	previousToDelete.next = previousToDelete.next.next
	l.length-- //handle the length
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 23}
	node2 := &node{data: 11}
	node3 := &node{data: 8}
	node4 := &node{data: 97}
	node5 := &node{data: 42}
	node6 := &node{data: 31}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)

	fmt.Println(mylist)
	mylist.printlistData()
	mylist.deleteWithValue(42)
	mylist.printlistData()
}
