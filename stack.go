package main

type Stack struct {
	topNode *StackNode
}

type StackNode struct {
	data     int
	nextNode *StackNode
}

func newStack(items ...int) Stack {
	s := Stack{}
	for _, item := range items {
		s.push(item)
	}
	return s
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

func (s *Stack) size() int {
	// returns the number of items in the data structure
	size := 0
	current := s.topNode
	for current != nil {
		size++
		current = current.nextNode
	}
	return size
}

func (s *Stack) top() StackNode {
	// returns the item that would be popped next
	nodeToRemove := s.topNode
	return *nodeToRemove
}

func (s *Stack) min() int {
	// returns the min integer data value in the stack
	min := s.topNode.data
	current := s.topNode
	for current != nil {
		if current.data < min {
			min = current.data
		}
		current = current.nextNode
	}
	return min

}

func (s *Stack) max() int {
	// returns the max integer data value in the stack
	max := s.topNode.data
	current := s.topNode
	for current != nil {
		if current.data > max {
			max = current.data
		}
		current = current.nextNode
	}
	return max
}
