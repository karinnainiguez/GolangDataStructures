package main

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	l := newList(3, 5, 7, 1)
	if reflect.TypeOf(l) != reflect.TypeOf(LinkedList{}) {
		t.Errorf("Expected a type of list, but instead received %v", reflect.TypeOf(l))
	}
}

func TestListCount(t *testing.T) {
	// Arrange
	l := newList(3, 5, 7, 1)

	// Act
	r1 := l.count()

	// Assert
	if r1 != 4 {
		t.Errorf("Expected a length of 4, but instead received %v", l.count())
	}
}

func TestListAdd(t *testing.T) {
	// Arrange
	l := newList(3, 5, 7, 1)
	originalCount := l.count()
	adding := 3
	expectedCount := originalCount + adding

	// Act
	l.add(3)
	l.add(8)
	l.add(2)

	r1 := l.count()

	// Assert
	if r1 != expectedCount {
		t.Errorf("Expected a length of %v, but instead received %v", expectedCount, l.count())
	}
}

func TestListIsFound(t *testing.T) {

	// Arrange
	newList := newList(3, 5, 7, 1)
	expectingToFind := 7
	notExpecting := 4

	// Act
	r1 := newList.isFound(expectingToFind)
	r2 := newList.isFound(notExpecting)

	// Assert
	if r1 == false {
		t.Errorf("Expect to find %v but returned false", expectingToFind)
	}

	if r2 == true {
		t.Errorf("Did not expect to find %v but returned true", notExpecting)
	}
}

func TestListMax(t *testing.T) {
	// Arrange
	l := newList(4, 2, 83, 5, 3, 7)
	expectedMax := 83

	// Act
	r1 := l.max()

	// Assert
	if r1 != expectedMax {
		t.Errorf("Expected max to be %v but received %v", expectedMax, r1)
	}
}

func TestListMin(t *testing.T) {
	// Arrange
	l := newList(8, 3, 87, 45, 1, 7)
	expectedMin := 1

	// Act
	r1 := l.min()

	// Assert
	if r1 != expectedMin {
		t.Errorf("Expected min to be %v but received %v", expectedMin, r1)
	}
}

func TestListValueAt(t *testing.T) {
	// Arrange
	l := newList(5, 7, 3, 9, 1)
	expectedAt0 := 1
	expectedAt2 := 3
	expectedAt3 := 7

	// Act
	r1 := l.valueAt(0)
	r2 := l.valueAt(2)
	r3 := l.valueAt(3)

	// Assert
	if r1 != expectedAt0 {
		t.Errorf("Expected %v but received %v", expectedAt0, r1)
	}

	if r2 != expectedAt2 {
		t.Errorf("Expected %v but received %v", expectedAt2, r2)
	}

	if r3 != expectedAt3 {
		t.Errorf("Expected %v but received %v", expectedAt3, r3)
	}
}

func TestListAscInsert(t *testing.T) {
	// Arrange
	l := newList()

	// Act
	l.ascInsert(78)
	l.ascInsert(23)
	l.ascInsert(5)
	l.ascInsert(86)
	l.ascInsert(1)
	l.ascInsert(79)

	// Assert
	if l.valueAt(0) != 1 {
		t.Errorf("Expected value at 0 to be 1, but received %v\n", l.valueAt(0))
	}

	if l.valueAt(1) != 5 {
		t.Errorf("Expected value at 0 to be 5, but received %v\n", l.valueAt(1))
	}

	if l.valueAt(2) != 23 {
		t.Errorf("Expected value at 0 to be 23, but received %v\n", l.valueAt(2))
	}

	if l.valueAt(3) != 78 {
		t.Errorf("Expected value at 0 to be 78, but received %v\n", l.valueAt(3))
	}

	if l.valueAt(4) != 79 {
		t.Errorf("Expected value at 0 to be 79, but received %v\n", l.valueAt(4))
	}
	if l.valueAt(5) != 86 {
		t.Errorf("Expected value at 0 to be 86, but received %v\n", l.valueAt(5))
	}

}

func TestListDeleteFirstVal(t *testing.T) {
	// Arrange
	l := newList(6, 3, 5, 2)

	// Act
	l.deleteFirstVal(5)

	// Assert
	r1 := l.count()
	r2 := l.isFound(5)

	if r1 != 3 {
		t.Errorf("Expected a list with length %v but received %v", 3, r1)
	}

	if r2 == true {
		t.Error("Did not expect to find 5 after deleting")
	}
}

func TestListReverseList(t *testing.T) {
	// Arrange
	l1 := newList(8, 4, 7, 3, 12)
	l2 := newList(12, 3, 7, 4, 8)

	// Act
	l1.reverseList()

	// Assert
	if l1.valueAt(3) != l2.valueAt(3) && l1.valueAt(4) != l2.valueAt(4) && l1.valueAt(5) != l2.valueAt(5) {
		t.Error("Expected list to be reversed, but it wasnt")
	}
}

func TestListReturnMiddleVal(t *testing.T) {
	// ODD
	// Arrange
	l := newList(6, 3, 9, 1, 8)
	expectedValue := 9

	// Act
	r1 := l.returnMiddleVal()

	// Assert
	if r1 != expectedValue {
		t.Errorf("Expected middle value of %v but instead received %v", expectedValue, r1)
	}

	// EVEN
	// Arrange
	l2 := newList(6, 3, 9, 1, 8, 2)
	newExpectedValue := 9

	// Act
	r2 := l2.returnMiddleVal()

	// Assert
	if r2 != newExpectedValue {
		t.Errorf("Expected middle value of %v but instead received %v", newExpectedValue, r2)
	}
}

func TestListValueAtFromEnd(t *testing.T) {
	// Arrange
	l := newList(5, 3, 8, 452, 572, 2, 12, 7)
	testingPosition := 4
	expectedValue := 452
	// Act
	r1 := l.valueAtFromEnd(testingPosition)

	// Assert
	if r1 != expectedValue {
		t.Errorf("Expecting value %v but instead received %v", expectedValue, r1)
	}
}

func TestListHasCycle(t *testing.T) {
	// Arrange
	l := newList(5, 4, 2, 3, 6)

	// Act
	r1 := l.hasCycle()

	// Assert
	if r1 == true {
		t.Error("Did not expect cycle, but returned true")
	}
}
