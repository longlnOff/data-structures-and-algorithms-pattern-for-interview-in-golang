package stack


type StackArray[T any] struct {
	elements []T
}

// Push adds a new element to the top of the stack
func (s *StackArray[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop removes the top element from the stack and returns it
func (s *StackArray[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

// Peek returns the top element of the stack without removing it
func (s *StackArray[T]) Peek() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty checks if the stack is empty
func (s *StackArray[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *StackArray[T]) Size() int {
	return len(s.elements)
}
