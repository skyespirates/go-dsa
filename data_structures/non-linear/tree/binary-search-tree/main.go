package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{}
}

func insertNode(node *Node, value int) {
	if value < node.value {
		if node.left == nil {
			node.left = &Node{value: value}
		} else {
			insertNode(node.left, value)
		}
	} else {
		if node.right == nil {
			node.right = &Node{value: value}
		} else {
			insertNode(node.right, value)
		}
	}
}

func (t *BST) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
		return
	}
	insertNode(t.root, value)

}

func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.value {
		return true
	}

	if value < node.value {
		return searchNode(node.left, value)
	}

	return searchNode(node.right, value)
}

func (t *BST) Search(value int) bool {
	return searchNode(t.root, value)
}

func main() {
	bst := NewBST()

	bst.Insert(5)
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(5)

	found := bst.Search(5)
	fmt.Printf("found? %v", found)
}
