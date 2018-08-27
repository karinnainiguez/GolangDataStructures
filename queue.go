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

func (q *Queue) dequeue() QueueNode {
	nodeToRemove := q.beginning
	q.beginning = q.beginning.nextNode
	return *nodeToRemove
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

func (q *Queue) size() int {
	// returns the number of items in the data structure
	size := 0
	current := q.beginning
	for current != nil {
		size++
		current = current.nextNode
	}
	return size
}

func (q *Queue) front() QueueNode {
	// returns the item that would be dequeued next
	nextNode := q.beginning
	return *nextNode
}

func (q *Queue) min() {
	// returns the min integer data value in the queue

}

func (q *Queue) max() {
	// returns the max integer data value in the queue

}
