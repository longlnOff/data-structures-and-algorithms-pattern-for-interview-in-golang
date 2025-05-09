package queue

// container/list: is used as linked-list
// fmt: used to return fmt.Errorf for the error part
import (
	"container/list"
	"errors"
)

// LQueue will be store the value into the list
type LQueue struct {
	elements *list.List
	size int
}

func NewQueueListArray() *LQueue {
	return &LQueue{
		elements: &list.List{},
		size: 0,
	}
}

// EnQueue it will be added new value into queue
func (q *LQueue)EnQueue(n any) {
	q.elements.PushBack(n)
	q.size += 1
}

// DeQueue it will be removed the first value that added into the list
func (q *LQueue)DeQueue() (any, error) {
	if q.IsEmptyQueue() {
		var emptyVal any
		return emptyVal, errors.New("queue is empty")
	} else {
		element := q.elements.Front()
		q.elements.Remove(element)
		q.size -= 1
		return element.Value, nil
	}
}

// FrontQueue return the Front value
func (q *LQueue)FrontQueue() (any, error) {
	if q.IsEmptyQueue() {
		var emptyVal any
		return emptyVal, errors.New("queue is empty")
	} else {
		return q.elements.Front().Value, nil
	}
}

// BackQueue return the Back value
func (q *LQueue)BackQueue() (any, error) {
	if q.IsEmptyQueue() {
		var emptyVal any
		return emptyVal, errors.New("queue is empty")
	} else {
		return q.elements.Back().Value, nil
	}
}

// LenQueue will return the length of the queue list
func (q *LQueue)LenQueue() int {
	return q.size
}

// IsEmptyQueue check our list is empty or not
func (q *LQueue)IsEmptyQueue() bool {
	return q.LenQueue() == 0
}
