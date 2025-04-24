package stack

type Node[T any] struct {
	Data T
	Next *Node[T]
}
type LinkedListStack[T any] struct {
	Top *Node[T]
	Length int
}

// Create new stack


// Push adds a new element to the top of the stack
func (lls *LinkedListStack[T]) Push(element T) {
	newNode := Node[T]{element, lls.Top}
	lls.Top = &newNode
	lls.Length += 1
}

// Pop removes the top element from the stack and returns it
func (lls *LinkedListStack[T]) Pop() T {
	if lls.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	data := lls.Top.Data
	lls.Top.Next = lls.Top
	lls.Length -= 1
	return data
}

// Peek returns the top element of the stack without removing it
func (lls *LinkedListStack[T]) Peek() T {
	if lls.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return lls.Top.Data
}

// IsEmpty checks if the stack is empty
func (lls *LinkedListStack[T]) IsEmpty() bool {
	return lls.Length == 0
}

// Size returns the number of elements in the stack
func (lls *LinkedListStack[T]) Size() int {
	return lls.Length
}
