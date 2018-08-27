package main

type Queue struct {
	beginning *QueueNode
	ending    *QueueNode
}

type QueueNode struct {
	data     int
	nextNode *QueueNode
}
