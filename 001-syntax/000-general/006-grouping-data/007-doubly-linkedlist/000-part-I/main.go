package main

import "fmt"

type node struct {
   value int
   prev *node
   next *node
}

func newNode(value int) *node{
   var node node
   node.next  = nil
   node.value = value
   node.prev  = nil
   return &node
}

func traverseDoublyLinkedlist(head * node){
   fmt.Printf("Doubly Linked List:\n ")
   temp := head
   for temp!= nil{
      fmt.Printf("%d ", temp.value)
	  fmt.Println("value: \taddress of[",&temp.value,"]")
      temp = temp.next
   }
}
func main(){

   head  := newNode(5)
   node2 := newNode(10)
   node3 := newNode(15)
   node4 := newNode(20)

   head.next  = node2
   node2.prev = head
   node2.next = node3
   node3.prev = node2
   node3.next = node4
   node4.prev = node3

   // Forward Traversal
   traverseDoublyLinkedlist(head)
}