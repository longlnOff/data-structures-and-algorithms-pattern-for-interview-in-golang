package stack

import (
	"container/list"
)

type Slist struct {
	Stack *list.List
}

// Push adds a new element to the top of the stack
func (s *Slist) Push(element any) {
	s.Stack.PushFront(element)
}
// Pop removes the top element from the stack and returns it
func (s *Slist) Pop() any {
	if s.IsEmpty() {
		var zeroValue any
		return zeroValue
	}
	element := s.Stack.Front()
	s.Stack.Remove(element)
	return element
}

// Peek returns the top element of the stack without removing it
func (s *Slist) Peek() any {
	if s.IsEmpty() {
		var zeroValue any
		return zeroValue
	}
	return s.Stack.Front()
}

// IsEmpty checks if the stack is empty
func (s *Slist) IsEmpty() bool {
	return s.Stack.Len() == 0
}

// Size returns the number of elements in the stack
func (s *Slist) Size() int {
	return s.Stack.Len()
}
