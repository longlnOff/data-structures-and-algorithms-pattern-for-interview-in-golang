package linkedlist

import (
	"errors"
	"fmt"
)


type Node[T any] struct {
	Data T
	Next *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{data, nil}
}

// Singly structure with length of the list and its head
type SinglyLinkedList[T any] struct {
	length int
	Head *Node[T]
}

// NewSingly returns a new instance of a linked list
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// AddAtBeg adds a new node with given value at the beginning of the list
func (ll *SinglyLinkedList[T]) AddAtBeg(data T) {
	// Create new node
	newNode := NewNode(data)
	// Next of new node = head
	newNode.Next = ll.Head
	// head = new node
	ll.Head = newNode
	// increase length
	ll.length += 1
}

// AddAtEnd adds a new node with given value at the end of the list
func (ll *SinglyLinkedList[T]) AddAtEnd(data T) {
	newNode := NewNode(data)
	if ll.Head == nil {
		ll.Head = newNode
		ll.length += 1
		return	
	}
	var cur *Node[T]
	for cur = ll.Head; cur.Next != nil; cur = cur.Next {}
	cur.Next = newNode
	ll.length += 1
}

// DelAtBeg deletes the node at the head(beginning) of the list
// and returns its value. Returns false if the list is empty
func (ll *SinglyLinkedList[T]) DelAtBeg() (T, bool) {
	if ll.length == 0 {
		var r T
		return r, false
	}
	data := ll.Head.Data
	ll.Head = ll.Head.Next
	ll.length -= 1
	return data, true
}

// DelAtEnd deletes the node at the tail(end) of the list
// and returns its value. Returns false if the list is empty
func (ll *SinglyLinkedList[T]) DelAtEnd() (T, bool) {
	if ll.length == 0 {
		var r T
		return r, false
	}
	// travel to the near end
	var cur *Node[T]
	for cur = ll.Head; cur.Next.Next != nil; cur = cur.Next {}
	data := cur.Next.Data
	cur.Next = cur.Next.Next
	ll.length -= 1
	return data, true
}

// DelByPos deletes the node at the middle based on position in the list
// and returns its value. Returns false if the list is empty or length is not more than given position
func (ll *SinglyLinkedList[T]) DelByPos(pos int) (T, bool) {
	// Check empty linked list
	if ll.length == 0 {
		var r T
		return r, false
	}
	// validate position < 1 or >= length
	if pos < 1 || pos > ll.length {
		var r T
		return r, false
	}
	// move to the before position and delete
	cur := ll.Head
	for index := 1; index < pos; index++ {
		cur = cur.Next
	}	
	data := cur.Data
	cur.Next = cur.Next.Next
	ll.length -= 1
	return data, true
}

// Count returns the current size of the list
func (ll *SinglyLinkedList[T]) Count() int {
	return ll.length
}

// Reverse reverses the list (inplace)
func (ll *SinglyLinkedList[T]) Reverse() {
	if ll.length == 0 || ll.length == 1 {
		return
	}

	cur := ll.Head
	var prev *Node[T]
	var next *Node[T]

	for cur != nil {
		next = cur.Next 
		cur.Next = prev
		prev = cur
		cur = next
	} 
	ll.Head = prev
}

// ReversePartition Reverse the linked list left the ath to the bth node (inplace)
func (ll *SinglyLinkedList[T]) ReversePartition(left int, right int) error {
	// check left and right
	if left < 1 || left > ll.length || right < 1 || right > ll.length || left >= right {
		return errors.New("invalid index to reverse")
	}
	tmpNode := &Node[T]{}
	tmpNode.Next = ll.Head
	prev := tmpNode
	// travel to left
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	// reverse left position to right position
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	} 
	ll.Head = tmpNode.Next
	return nil
}

// Display prints out the elements of the list
func (ll *SinglyLinkedList[T]) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " --> ")
	}

	fmt.Println("End")
}
