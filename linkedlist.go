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

// 7. Insert new node with specific input data value, assuming
// list is sorted in ascending order.
func (l *LinkedList) insertAsc(data int) {
	if l.head == nil || l.head.data > data {
		newNode := Node{data: data, nextNode: l.head}
		l.head = &newNode
	} else {
		current := l.head
		for current.nextNode.data < data {
			current = current.nextNode
		}
		newNode := Node{data: data, nextNode: current.nextNode}
		current.nextNode = &newNode
	}
}

// 8. Print the value of each node in list separated by space.
func (l *LinkedList) printValues() {
	
}

// 9. Delete first node found with specified input data value.
func (l *LinkedList) () {}

// 10.  Reverse the list.  Nodes should be preserved, not
// just their values.
func (l *LinkedList) () {}

// 11. Return the value at the middle node.  If node count
// is even, pick one of the two middle nodes to return.
func (l *LinkedList) () {}

// 12. Find the nth node from the end of list and return
// value.  Assuming indexing starts at 0 while count to n
func (l *LinkedList) () {}

// 13. Check if hte linked list has a cycle.  A cycle
// exists if any node in the linked list links to a node
// already visited.
func (l *LinkedList) () {}



/// MAIN FUNCTION: 
func main() {
	fmt.Println("Working")
}
