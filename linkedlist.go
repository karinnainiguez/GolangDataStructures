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

// 3. return max value in list (return data, not node)
func (l *LinkedList) max() int {
	max := l.head.data
	current := l.head
	for current != nil {
		if current.data > max {
			max = current.data
		}
	}
	return max
}

// 4. return min value in list (return data, not node)
func (l *LinkedList) min() int {
	min := l.head.data
	current := l.head
	for current != nil {
		if current.data < min {
			min = current.data
		}
	}
	return min
}

// 5. return count of nodes in list
func (l *LinkedList) count() int {
	count := 0
	current := l.head
	for current != nil {
		count++
	}
	return count
}

// 6. Return the value of the nth node from the beginning.
// n is the input to the method. Assume indexing starts at
// 0 while counting to n.
func (l *LinkedList) valueAt(n int) int {
	count := 0
	current := l.head
	for current != nil {
		count++
		if count == n {
			return current.data
		}
	}
	return 0
}

func main() {
	fmt.Println("Working")
}
