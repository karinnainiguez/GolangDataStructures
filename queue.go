package main

type Queue struct {
	beginning *QueueNode
	ending    *QueueNode
}

type QueueNode struct {
	data     int
	nextNode *QueueNode
}

func (q *Queue) enqueue() {

}

func (q *Queue) dequeue() {

}

func (q *Queue) Empty() bool {
	// returns true is the data structure is empty,
	// false otherwise
	return false
}

func (q *Queue) Size() {

}

func (q *Queue) Front() {

}

func (q *Queue) Min() {

}

func (q *Queue) Max() {

}
