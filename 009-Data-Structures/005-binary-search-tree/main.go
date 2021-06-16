package main

import "fmt"

//node represents the components of a binary search tree..
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//insert will add a node to the tree :: we need to manipulate the receiver so we need a pointer
func (n *Node) Insert(k int) {
	//r := "right: "
	//l := "left: "
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
		// End of the inner check
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//search
//TODO

func main() {
	tree := &Node{Key: 1000}
	tree.Insert(999)
	fmt.Println(tree)
}
