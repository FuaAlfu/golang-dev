package main

import "fmt"

//node represents the components of a binary search tree..
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//count nodes
var count int

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
/*
will take in a key value
and return true if there is a node with that value
*/
//TODO
func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		//if it nill then its means that, the end of tree..
		//we didn't meat a match
		return false
	}
	if n.Key < k {
		//move right :: if right will repeat the method Search
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left :: if left will repeat the method Search
		return n.Left.Search(k)
	}

	return true
}

func main() {
	/*
		ROOT > PARENT > CHILDREN > LEAVES
	*/
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(999)
	tree.Insert(40)
	tree.Insert(137)
	tree.Insert(3)
	tree.Insert(312)
	tree.Insert(70)
	tree.Insert(31)
	fmt.Println(tree.Search(40))
	println(count)
}
