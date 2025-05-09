package queue

import "errors"

type QueueArray[T any] struct {
	elements []T
	size     int
}

func NewQueueArray[T any]() *QueueArray[T] {
	return &QueueArray[T]{
		elements: []T{},
		size:     0,
	}
}

// EnQueue it will be added new value into queue
func (q *QueueArray[T])EnQueue(n T) {
	q.elements = append(q.elements, n)
	q.size += 1
}

// DeQueue it will be removed the first value that added into the list
func (q *QueueArray[T])DeQueue() (T, error) {
	if q.IsEmptyQueue() {
		var emptyVal T
		return emptyVal, errors.New("queue is empty")
	} else {
		q.size -= 1
		retVal := q.elements[0]
		q.elements = q.elements[1:]
		return retVal, nil
	}
}

// FrontQueue return the Front value
func (q *QueueArray[T])FrontQueue() (T, error) {
	if q.IsEmptyQueue() {
		var emptyVal T
		return emptyVal, errors.New("queue is empty")
	} else {
		return q.elements[0], nil
	}
}

// BackQueue return the Back value
func (q *QueueArray[T])BackQueue() (T, error) {
	if q.IsEmptyQueue() {
		var emptyVal T
		return emptyVal, errors.New("queue is empty")
	} else {
		return q.elements[q.size-1], nil
	}
}

// LenQueue will return the length of the queue list
func (q *QueueArray[T])LenQueue() int {
	return q.size
}

// IsEmptyQueue check our list is empty or not
func (q *QueueArray[T])IsEmptyQueue() bool {
	return q.LenQueue() == 0
}
