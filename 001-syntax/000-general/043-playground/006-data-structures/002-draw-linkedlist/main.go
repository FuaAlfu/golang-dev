package main

import "fmt"

type node struct {
   value int
   next *node
}

func newNode(value int) *node{
   var n node
   n.value = value
   n.next = nil
   return &n
}

func traverseLinkedList(head *node){
   fmt.Printf("Linked List: ")
   temp := head
   for temp != nil {
      fmt.Printf("%d -> ", temp.value)
      temp = temp.next
   }
   fmt.Printf("NULL")
}
func main(){
   head := newNode(300)
   head.next = newNode(289)
   head.next.next = newNode(73)
   head.next.next.next = newNode(8)
   traverseLinkedList(head)
}
