package main

type Stack struct {
	top *StackNode
}

type StackNode struct {
	data     int
	nextNode *StackNode
}

func (s *Stack) push(data int) {

}

func (s *Stack) pop() StackNode {
	nodeToRemove := s.top
	s.top = s.top.nextNode
	return *nodeToRemove
}
