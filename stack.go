package main

type Stack struct {
	topNode *StackNode
}

type StackNode struct {
	data     int
	nextNode *StackNode
}

func (s *Stack) push(data int) {
	newNode := StackNode{data: data, nextNode: s.topNode}
	s.topNode = &newNode
}

func (s *Stack) pop() StackNode {
	nodeToRemove := s.topNode
	s.topNode = s.topNode.nextNode
	return *nodeToRemove
}

func (s *Stack) empty() bool {
	// returns true is the data structure is empty
	if s.topNode == nil {
		return true
	}
	return false

}

func (s *Stack) size() {
	// returns the number of items in the data structure

}

func (s *Stack) top() {
	// returns the item that would be popped next

}

func (s *Stack) min() {
	// returns the min integer data value in the stack
}

func (s *Stack) max() {
	// returns the max integer data value in the stack

}
