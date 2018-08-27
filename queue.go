package main

type Queue struct {
	beginning *QueueNode
	ending    *QueueNode
}

type QueueNode struct {
	data     int
	nextNode *QueueNode
}

func (q *Queue) enqueue(data int) {
	if q.beginning == nil {
		newNode := QueueNode{data: data}
		q.beginning = &newNode
		q.ending = &newNode
	} else {
		newNode := QueueNode{data: data}
		q.ending.nextNode = &newNode
		q.ending = &newNode
	}
}

func (q *Queue) dequeue() {

}

func (q *Queue) empty() bool {
	// returns true is the data structure is empty,
	// false otherwise
	if q.beginning == nil {
		return true
	} else {
		return false
	}
}

func (q *Queue) size() {

}

func (q *Queue) front() {

}

func (q *Queue) min() {

}

func (q *Queue) max() {

}
