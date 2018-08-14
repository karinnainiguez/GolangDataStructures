package main

import (
	"fmt"
)

type Node struct {
	data     int
	nextNode *Node
}

type LinkedList struct {
	head *Node
}

// 1. Add new node to list based on input data
func (l *LinkedList) add(inputData int) {
	newNode := Node{data: inputData, nextNode: l.head}
	l.head = &newNode
}

// 2. Check if list contains node with specified input value (boolean)
func (l *LinkedList) isFound(value int) bool {
	current := l.head
	for current != nil {
		if current.data == value {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Working")
}
