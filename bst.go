package main

type binarySearchTree struct {
	rootNode *bstNode
}

type bstNode struct {
	data       int
	rightChild *bstNode
	leftChild  *bstNode
}

// Create a new binary search tree
func newBST(items ...int) binarySearchTree {
	bst := binarySearchTree{}
	for _, item := range items {
		bst.add(item)
	}
	return bst
}

// count nodes in tree
func (bst *binarySearchTree) count() int {
	if bst.rootNode == nil {
		return 0
	}
	return countHelper(bst.rootNode)
}

func countHelper(parent *bstNode) int {
	if parent == nil {
		return 0
	}
	return 1 + countHelper(parent.leftChild) + countHelper(parent.rightChild)
}

// insert given integer value in the tree
func (bst *binarySearchTree) add(item int) {
	newNode := bstNode{data: item}

	if bst.rootNode == nil {
		bst.rootNode = &newNode
	} else {
		addHelper(bst.rootNode, newNode)
	}
}

func addHelper(current *bstNode, adding bstNode) {
	if adding.data < current.data {
		// either assign to left child,
		// or recursively call on left child
		if current.leftChild == nil {
			current.leftChild = &adding
		} else {
			addHelper(current.leftChild, adding)
		}

	} else {
		// either assign to right child,
		// or recursively call on right child
		if current.rightChild == nil {
			current.rightChild = &adding
		} else {
			addHelper(current.rightChild, adding)
		}
	}
}

// search for a given integer value in the tree (return bool)
func (bst *binarySearchTree) search(num int) bool {

	if bst.rootNode == nil {
		return false
	}
	return searchHelper(bst.rootNode, num)

}

// recursive search helper
func searchHelper(node *bstNode, num int) bool {
	if node.data == num {
		return true
	}
	if node.leftChild != nil && searchHelper(node.leftChild, num) == true {
		return true
	}
	if node.rightChild != nil {
		return searchHelper(node.rightChild, num)
	}
	return false
}

// compute the height of the tree
func (bst *binarySearchTree) height() int {
	return heightHelper(bst.rootNode)
}

// recursive height helper
func heightHelper(node *bstNode) int {
	if node == nil {
		return 0
	}
	left := heightHelper(node.leftChild)
	right := heightHelper(node.rightChild)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

// print values in BST in PRE-ORDER
func (bst *binarySearchTree) preOrder() []int {
	r := []int{}
	preOrderHelper(bst.rootNode, &r)
	return r
}
func preOrderHelper(node *bstNode, slice *[]int) {
	if node == nil {
		return
	}
	*slice = append(*slice, node.data)
	preOrderHelper(node.leftChild, slice)
	preOrderHelper(node.rightChild, slice)
}

// print values in BST IN-ORDER
func (bst *binarySearchTree) inOrder() []int {
	r := []int{}
	inOrderHelper(bst.rootNode, &r)
	return r
}
func inOrderHelper(node *bstNode, slice *[]int) {
	if node == nil {
		return
	}
	inOrderHelper(node.leftChild, slice)
	*slice = append(*slice, node.data)
	inOrderHelper(node.rightChild, slice)
}

// print values in BST in POST-ORDER
func (bst *binarySearchTree) postOrder() []int {
	r := []int{}
	postOrderHelper(bst.rootNode, &r)
	return r
}
func postOrderHelper(node *bstNode, slice *[]int) {
	if node == nil {
		return
	}
	postOrderHelper(node.leftChild, slice)
	postOrderHelper(node.rightChild, slice)
	*slice = append(*slice, node.data)
}
