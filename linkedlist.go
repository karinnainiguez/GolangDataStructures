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

func newList(items ...int) LinkedList {
	l := LinkedList{}
	for _, item := range items {
		l.add(item)
	}
	return l
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
		current = current.nextNode
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
		current = current.nextNode
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
		current = current.nextNode
	}
	return min
}

// 5. return count of nodes in list
func (l *LinkedList) count() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.nextNode
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
		if count == n {
			return current.data
		}
		count++
		current = current.nextNode
	}
	return 0
}

// 7. Insert new node with specific input data value, assuming
// list is sorted in ascending order.
func (l *LinkedList) ascInsert(inputData int) {
	newNode := Node{data: inputData}

	// if list is empty, just insert node as head
	if l.head == nil {
		l.head = &newNode
	} else if l.head.data > inputData {
		// if the head is bigger than the input data, insert node as head's nextNode
		newNode.nextNode = l.head
		l.head = &newNode
	} else {

		current := l.head

		if current.nextNode == nil {
			current.nextNode = &newNode
		} else {
			// otherwise, iterate through list, and stop when the current's next node is bigger
			// than the input data, or when you've reached the end of the list.
			for current.nextNode != nil && current.nextNode.data < inputData {
				current = current.nextNode
			}

			// assign next node (which will be either bigger than data, or nil) as the newNode's next
			// and current's next node as newNode
			newNode.nextNode = current.nextNode
			current.nextNode = &newNode

		}
	}
}

// 8. Print the value of each node in list separated by space.
func (l *LinkedList) printValues() {
	current := l.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.nextNode
	}
}

// 9. Delete first node found with specified input data value.
func (l *LinkedList) deleteFirstVal(data int) {
	current := l.head
	for current.nextNode != nil && current.nextNode.data != data {
		current = current.nextNode
	}
	if current.nextNode != nil {
		current.nextNode = current.nextNode.nextNode
	}
}

// 10.  Reverse the list.  Nodes should be preserved, not
// just their values.
func (l *LinkedList) reverseList() {

	var previous *Node = nil
	current := l.head
	var next *Node = nil
	for current != nil {
		next = current.nextNode
		current.nextNode = previous

		if next == nil {
			l.head = current
		}
		previous = current
		current = next
	}

}

// 11. Return the value at the middle node.  If node count
// is even, pick one of the two middle nodes to return.
func (l *LinkedList) returnMiddleVal() int {
	nodeCount := l.count()
	middle := nodeCount / 2
	current := l.head
	for i := 0; i < middle; i++ {
		current = current.nextNode
	}
	return current.data
}

// 12. Find the nth node from the end of list and return
// value.  Assuming indexing starts at 0 while count to n
func (l *LinkedList) valueAtFromEnd(n int) int {
	count := l.count()
	goal := count - n
	if goal < 0 {
		return 0
	}
	current := l.head
	for i := 0; i < goal; i++ {
		current = current.nextNode
	}
	return current.data
}

// 13. Check if the linked list has a cycle.  A cycle
// exists if any node in the linked list links to a node
// already visited.
func (l *LinkedList) hasCycle() bool {
	seen := []Node{}
	current := l.head
	for current != nil {
		for _, node := range seen {
			if *current == node {
				return true
			}
			seen = append(seen, node)
		}
		current = current.nextNode
	}
	return false
}
