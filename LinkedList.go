package udemyproblems

import "fmt"

/*

-- AddatFront      -- O(1)
-- AddatEnd        -- O(1)

-- RemoveFromFront -- O(1)
-- RemoveFromEnd   -- O(N)

*/
//Node of a singly linked list
type Node struct {
	value int
	next  *Node
}

//LinkedList defines a singly linked list
type LinkedList struct {
	head *Node
	tail *Node
}

//IsLinkedListEmpty returns true if the given linked list is empty
func (L *LinkedList) IsLinkedListEmpty() bool {

	isempty := false

	//condition for emply list is both head and tail are nil
	if L.head == nil && L.tail == nil {
		isempty = true
	}

	return isempty
}

//PrintLinkedList prints the values in the linked list
func PrintLinkedList(L *LinkedList) {

	if L.IsLinkedListEmpty() {
		fmt.Println("The given Linked List has no elements")
	} else {
		currentnode := L.head
		for currentnode != nil {
			fmt.Println(currentnode.value)
			currentnode = currentnode.next
		}

	}
}

//AddatEnd adds a given value to the end of the list
func (L *LinkedList) AddatEnd(Value int) {

	//setup newnode
	newnode := Node{}
	newnode.value = Value

	if L.IsLinkedListEmpty() {
		//set both head and tail
		L.head = &newnode
	} else {
		L.tail.next = &newnode
	}

	//move the tail to point to newnode
	L.tail = &newnode

}

//RemoveFromFront removes the node pointed by the head
func (L *LinkedList) RemoveFromFront() {

	if !(L.IsLinkedListEmpty()) {

		currentnode := L.head
		//move head forward
		L.head = L.head.next
		currentnode.next = nil
		
		//fmt.Println("Removed node with value = ", currentnode.value)

		if L.head == nil {
			L.tail = nil
		}
	}
}
