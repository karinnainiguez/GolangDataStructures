package main

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	// Arrange
	s := newStack(4, 3, 7)

	// Act
	r1 := reflect.TypeOf(s)

	// Assert
	if r1 != reflect.TypeOf(Stack{}) {
		t.Errorf("Expected to return Stack, but instead returned %v", r1)
	}
}

func TestStackSize(t *testing.T) {
	// Arrange
	s1 := newStack(3, 52, 7)
	expectedSize1 := 3
	s2 := newStack(7, 2, 4, 7, 93, 2)
	expectedSize2 := 6
	s3 := newStack(9)
	expectedSize3 := 1

	// Act
	r1 := s1.size()
	r2 := s2.size()
	r3 := s3.size()

	// Assert
	if r1 != expectedSize1 {
		t.Errorf("Expected a stack with size %v but received %v\n", expectedSize1, r1)
	}

	if r2 != expectedSize2 {
		t.Errorf("Expected a stack with size %v but received %v\n", expectedSize2, r2)
	}

	if r3 != expectedSize3 {
		t.Errorf("Expected a stack with size %v but received %v\n", expectedSize3, r3)
	}
}

func TestStackPush(t *testing.T) {
	// Arrange
	s := newStack(5, 3, 7)
	adding := 3
	originalSize := s.size()

	// Act
	s.push(6)
	s.push(9)
	s.push(1)

	r1 := s.size()

	// Assert
	if r1 != originalSize+adding {
		t.Errorf("Expected stack of size %v but instead have a size of %v", originalSize+adding, r1)
	}
}

func TestStackPop(t *testing.T) {
	// Arrange
	s := newStack(3, 4, 5, 6, 7)
	removing := 2
	originalSize := s.size()

	// Act
	s.pop()
	s.pop()

	r1 := s.size()

	// Assert
	if r1 != originalSize-removing {
		t.Errorf("Expected stack of size %v but instead has size of %v", originalSize-removing, r1)
	}
}

func TestStackEmpty(t *testing.T) {
	// Arrange
	s1 := newStack()
	s2 := newStack(4, 5, 2)
	s3 := newStack(2)

	// Act
	s3.pop()

	r1 := s1.empty()
	r2 := s2.empty()
	r3 := s3.empty()

	// Assert
	if r1 == false {
		t.Error("Expected stack to be empty but received false")
	}

	if r2 == true {
		t.Error("Expected stack to NOT be empty but received true")
	}

	if r3 == false {
		t.Error("Expected stack to be empty but received false")
	}
}

func TestStackTop(t *testing.T) {
	// Arrange
	s := newStack(5, 2, 6, 2, 9)
	originalSize := s.size()

	// Act
	r1 := s.top()
	r2 := s.size()

	// Assert
	if r1.data != 9 {
		t.Errorf("Expected next item to be 9 but it was %v", r1)
	}

	if r2 != originalSize {
		t.Errorf("Did not expect stack size to change original size: %v new size: %v", originalSize, r2)
	}

}

func TestStackMin(t *testing.T) {
	// Arrange
	s := newStack(5, 3, 8, 95, 2, 81)
	expectedMin := 2

	// Act
	r1 := s.min()

	// Assert
	if r1 != expectedMin {
		t.Errorf("Expected a min of %v but instead returned %v", expectedMin, r1)
	}

}

func TestStackMax(t *testing.T) {
	// Arrange
	s := newStack(5, 3, 8, 95, 2, 81)
	expectedMax := 95

	// Act
	r1 := s.max()

	// Assert
	if r1 != expectedMax {
		t.Errorf("Expected a max of %v but instead returned %v", expectedMax, r1)
	}
}
