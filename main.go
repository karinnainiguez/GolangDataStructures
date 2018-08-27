package main

import "fmt"

/// MAIN FUNCTION:
func main() {

	fmt.Println("\nCreating first list...")
	list := LinkedList{}
	list.printValues()

	fmt.Println("\n (insert) Adding three values...")
	list.add(8)
	list.add(3)
	list.add(15)
	list.printValues()

	find8 := list.isFound(8)
	fmt.Printf("\n(search) Check to see if value 8 is included: %v\n", find8)

	max := list.max()
	fmt.Printf("\n(find max) Max in current list should be 15.  Result: %v\n", max)

	min := list.min()
	fmt.Printf("\n(find min) Min in current list should be 3.  Result: %v\n", min)

	length := list.count()
	fmt.Printf("\n(length) Current node count should be 3.  Result: %v\n", length)

	nthBeginning := list.valueAt(2)
	fmt.Printf("\n(find nth from beginning) currentList.valueAt(2) should return 8.  Result: %v\n", nthBeginning)

	fmt.Println("\n(visit) Printing each value in linked list...")
	list.printValues()
	fmt.Printf("\n")

	fmt.Println("\n(delete) going to delete node with value 8.  Printing list again...")
	list.deleteFirstVal(8)
	list.printValues()
	fmt.Printf("\n")

	fmt.Printf("\nCreating new list to try to reverse it... \n")
	newList := LinkedList{}
	newList.add(8)
	newList.add(6)
	newList.add(10)
	newList.add(7)
	newList.add(5)
	newList.add(9)
	fmt.Println("original list: ")
	newList.printValues()
	fmt.Println("\nReversed: ")
	newList.reverseList()
	newList.printValues()
	fmt.Printf("\n\n")

	fmt.Printf("\n(ascending list) Creating ascending list...\n")
	ascList := LinkedList{}
	ascList.ascInsert(59)
	ascList.ascInsert(45)
	ascList.ascInsert(64)
	ascList.ascInsert(24)
	ascList.ascInsert(81)
	ascList.ascInsert(85)
	ascList.ascInsert(56)
	ascList.ascInsert(87)
	ascList.ascInsert(101)
	ascList.ascInsert(99)
	ascList.printValues()

	fmt.Printf("\n\n\n")
}
