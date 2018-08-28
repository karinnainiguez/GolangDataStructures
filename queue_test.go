package main

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	// Arange
	q := newQueue(5, 3, 4)

	// Act
	r1 := reflect.TypeOf(q)

	// Assert
	if r1 != reflect.TypeOf(Queue{}) {
		t.Errorf("Expected to return Queue, but instead returned %v\n", r1)
	}

}

func TestQueueSize(t *testing.T) {
	// Arange
	q1 := newQueue(2, 4, 65, 8)
	expectedSize1 := 4
	q2 := newQueue(2, 4, 65, 8, 68, 96, 34)
	expectedSize2 := 7
	q3 := newQueue(2)
	expectedSize3 := 1

	// Act
	r1 := q1.size()
	r2 := q2.size()
	r3 := q3.size()

	// Assert
	if r1 != expectedSize1 {
		t.Errorf("Expected size of %v but returned size of %v \n", expectedSize1, r1)
	}

	if r2 != expectedSize2 {
		t.Errorf("Expected size of %v but returned size of %v \n", expectedSize2, r2)
	}

	if r3 != expectedSize3 {
		t.Errorf("Expected size of %v but returned size of %v \n", expectedSize3, r3)
	}

}

func TestQueueEnqueue(t *testing.T) {
	// Arange
	q := newQueue(6, 9, 3)
	adding := 2
	originalSize := q.size()

	// Act
	q.enqueue(13)
	q.enqueue(7)

	r1 := q.size()

	// Assert
	if r1 != originalSize+adding {
		t.Errorf("Expected queue to have size of %v but has size of %v", originalSize+adding, r1)
	}

}

func TestQueueDequeue(t *testing.T) {
	// Arange
	q := newQueue(5, 4, 7, 6, 3)
	removing := 2
	originalSize := q.size()

	// Act
	q.dequeue()
	q.dequeue()

	r1 := q.size()

	// Assert
	if r1 != originalSize-removing {
		t.Errorf("Expecting queue to have size of %v but has size of %v", originalSize-removing, r1)
	}

}

func TestQueueEmpty(t *testing.T) {
	// Arange
	q1 := newQueue()
	q2 := newQueue(7, 3, 5)
	q3 := newQueue(5)

	// Act
	q3.dequeue()

	r1 := q1.empty()
	r2 := q2.empty()
	r3 := q3.empty()

	// Assert

	if r1 == false {
		t.Error("Expected queue to be empty, but returned false")
	}

	if r2 == true {
		t.Error("Expected queue to NOT be empty, but returned true")
	}

	if r3 == false {
		t.Error("Expected queue to be empty, but returned false")
	}

}

func TestQueueFront(t *testing.T) {
	// Arange

	// Act

	// Assert

}

func TestQueueMin(t *testing.T) {
	// Arange
	q := newQueue(9, 4, 7, 2, 5, 3)
	expectedMin := 2

	// Act
	r1 := q.min()

	// Assert
	if r1 != expectedMin {
		t.Errorf("Expecting queue to have size of %v but returned size of %v \n", expectedMin, r1)
	}

}

func TestQueueMax(t *testing.T) {
	// Arange
	q := newQueue(5, 3, 6, 2, 8, 1, 4)
	expectedMax := 8

	// Act
	r1 := q.max()

	// Assert
	if r1 != expectedMax {
		t.Errorf("Expected max to be %v but returend %v\n", expectedMax, r1)
	}

}
