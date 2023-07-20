package main

import "fmt"

type(
	node struct{
		data int
		next *node
	}
	linkedlist struct{
		head *node
	}
)

var mylist linkedlist

func(list *linkedlist) add(data int){
	newNode := &node(data, nil)
	if list.head == nil{
		list.head = newNode
	}else{
		//start at the begining..
		currentNode := list.head

		//loop through the list to find the end of list
		for currentNode != nil{
			currentNode = currentNode.next
		}
	    currentNode = newNode
	}
}

func(list *linkedlist) printList(){
	currentNode := list.head
	for currentNode != nil{
		fmt.Printf("current node address: %p \t node data %d \t next node: %p \n",currentNode, currentNode.data, currentNode.next)
		currentNode = currentNode.next
	}
}

func main(){
	fmt.Println("running..")
	mylist.add(1)
	mylist.add(2)
	mylist.add(3)
	mylist.printList()
}