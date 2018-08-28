package main

import (
	"reflect"
	"testing"
)

func TestNewBST(t *testing.T) {
	bst := newBST(4, 5, 6, 2)
	if reflect.TypeOf(bst) != reflect.TypeOf(binarySearchTree{}) {
		t.Errorf("Expected type of binary search tree, but returned %v\n", reflect.TypeOf(bst))
	}
}

func TestBSTCount(t *testing.T) {

	// Arrange
	bst1 := newBST(5, 2, 7, 3)
	expectedCount1 := 4

	bst2 := newBST(2)
	expectedCount2 := 1

	bst3 := newBST(7, 4, 2, 9, 12, 53, 8, 3)
	expectedCount3 := 8

	// Act
	r1 := bst1.count()
	r2 := bst2.count()
	r3 := bst3.count()

	// Assert
	if r1 != expectedCount1 {
		t.Errorf("Expected count %v but returned %v\n", expectedCount1, r1)
	}
	if r2 != expectedCount2 {
		t.Errorf("Expected count %v but returned %v\n", expectedCount2, r2)
	}
	if r3 != expectedCount3 {
		t.Errorf("Expected count %v but returned %v\n", expectedCount3, r3)
	}

}

func TestBSTAdd(t *testing.T) {

	// Arrange
	bst := newBST(4, 23, 5)
	originalCount := bst.count()
	adding := 4

	// Act
	bst.add(3)
	bst.add(37)
	bst.add(19)
	bst.add(10)

	r := bst.count()

	// Assert
	if r != originalCount+adding {
		t.Errorf("Expecting count of %v but ended up with count of %v\n", originalCount+adding, r)
	}
}

func TestBSTSearch(t *testing.T) {

	// Arrange
	bst := newBST(4, 7, 3, 8, 10, 2)
	// Act
	r1 := bst.search(7)
	r2 := bst.search(12)
	r3 := bst.search(10)
	r4 := bst.search(23)
	// Assert
	if r1 != true {
		t.Error("Expected 7 to be found")
	}
	if r2 != false {
		t.Error("Expected 12 to NOT be found")
	}
	if r3 != true {
		t.Error("Expected 10 to be found")
	}
	if r4 != false {
		t.Error("Expected 23 to Not be found")
	}
}

func TestBSTHeight(t *testing.T) {

	// Arrange
	bst1 := newBST(5, 3, 1, 4, 7, 12, 6, 9, 23, 11, 32)
	// Act
	r1 := bst1.height()
	expected1 := 5
	// Assert
	if r1 != expected1 {
		t.Errorf("Expected height of %v but returned height of %v", expected1, r1)
	}

	// SECOND TREE
	// Arrange
	bst2 := newBST(90, 34, 102, 23, 40, 100)
	// Act
	r2 := bst2.height()
	expected2 := 3
	// Assert
	if r2 != expected2 {
		t.Errorf("Expected height of %v but returned height of %v", expected2, r2)
	}

	// EMPTY TREE
	// Arrange
	bst3 := newBST()
	// Act
	r3 := bst3.height()
	expected3 := 0
	// Assert
	if r3 != expected3 {
		t.Errorf("Expected height of %v but returned height of %v", expected3, r3)
	}
}

func TestBSTPreOrder(t *testing.T) {

	// Arrange
	bst := newBST(6, 3, 8, 5, 9, 1, 7, 13)
	// Act
	r := bst.preOrder()
	expected := []int{6, 3, 1, 5, 8, 7, 9, 13}
	// Assert
	if reflect.DeepEqual(r, expected) != true {
		t.Errorf("Expected preorder of %v but returned %v", expected, r)
	}
}

func TestBSTInOrder(t *testing.T) {

	// Arrange
	bst := newBST(6, 3, 8, 5, 9, 1, 7, 13)
	// Act
	r := bst.inOrder()
	expected := []int{1, 3, 5, 6, 7, 8, 9, 13}
	// Assert
	if reflect.DeepEqual(r, expected) != true {
		t.Errorf("Expected inorder of %v but returned %v", expected, r)
	}
}

func TestBSTPostOrder(t *testing.T) {

	// Arrange
	bst := newBST(6, 3, 8, 5, 9, 1, 7, 13)
	// Act
	r := bst.postOrder()
	expected := []int{1, 5, 3, 7, 13, 9, 8, 6}
	// Assert
	if reflect.DeepEqual(r, expected) != true {
		t.Errorf("Expected postorder of %v but returned %v", expected, r)
	}
}
