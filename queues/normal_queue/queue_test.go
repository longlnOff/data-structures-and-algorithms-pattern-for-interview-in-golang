package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	// Test array queue
	t.Run("Test Array Queue", func(t *testing.T) {

		t.Run("Test EnQueue", func(t *testing.T) {
			newQueue := NewQueueArray[int]()
			newQueue.EnQueue(2)
			newQueue.EnQueue(3)
			newQueue.EnQueue(4)
			newQueue.EnQueue(45)
			
			// Validate front and back queue
			front, err := newQueue.FrontQueue()
			if err != nil {
				t.Errorf("Run FrontQueue Error")
			}
			back, err := newQueue.BackQueue()
			if err != nil {
				t.Errorf("Run FrontQueue Error")
			}
			if front != 2 && back != 45 {
				t.Errorf("Test EnQueue is wrong the result must be %v and %v but got %v and %v", 2, 45, front, back)
			}

		})

		t.Run("Test DeQueue", func(t *testing.T) {
			newQueue := NewQueueArray[int]()
			newQueue.EnQueue(2)
			newQueue.EnQueue(3)
			newQueue.EnQueue(4)

			newQueue.DeQueue()
			if val, err := newQueue.DeQueue(); val != 3 && err == nil {
				t.Errorf("Test DeQueue is wrong the result must be %v but got %v", 3, val)
			}

		})

		t.Run("Test Queue isEmpty", func(t *testing.T) {
			newQueue := NewQueueArray[int]()
			if newQueue.IsEmptyQueue() != true {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", true, newQueue.IsEmptyQueue())
			}

			newQueue.EnQueue(3)
			newQueue.EnQueue(4)

			if newQueue.IsEmptyQueue() != false {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", false, newQueue.IsEmptyQueue())
			}
		})

		t.Run("Test Queue Length", func(t *testing.T) {
			newQueue := NewQueueArray[int]()
			if newQueue.LenQueue() != 0 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 0, newQueue.LenQueue())
			}

			newQueue.EnQueue(3)
			newQueue.EnQueue(4)
			newQueue.DeQueue()
			newQueue.EnQueue(22)
			newQueue.EnQueue(99)
			newQueue.DeQueue()
			newQueue.DeQueue()

			if newQueue.LenQueue() != 1 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 1, newQueue.LenQueue())
			}

		})
	})


	t.Run("Test List Array Queue", func(t *testing.T) {

		t.Run("Test EnQueue", func(t *testing.T) {
			newQueue := NewQueueListArray()
			newQueue.EnQueue(2)
			newQueue.EnQueue(3)
			newQueue.EnQueue(4)
			newQueue.EnQueue(45)
			
			// Validate front and back queue
			front, err := newQueue.FrontQueue()
			if err != nil {
				t.Errorf("Run FrontQueue Error")
			}
			back, err := newQueue.BackQueue()
			if err != nil {
				t.Errorf("Run FrontQueue Error")
			}
			if front != 2 && back != 45 {
				t.Errorf("Test EnQueue is wrong the result must be %v and %v but got %v and %v", 2, 45, front, back)
			}

		})

		t.Run("Test DeQueue", func(t *testing.T) {
			newQueue := NewQueueArray[int]()
			newQueue.EnQueue(2)
			newQueue.EnQueue(3)
			newQueue.EnQueue(4)

			newQueue.DeQueue()
			if val, err := newQueue.DeQueue(); val != 3 && err == nil {
				t.Errorf("Test DeQueue is wrong the result must be %v but got %v", 3, val)
			}

		})

		t.Run("Test Queue isEmpty", func(t *testing.T) {
			newQueue := NewQueueListArray()
			if newQueue.IsEmptyQueue() != true {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", true, newQueue.IsEmptyQueue())
			}

			newQueue.EnQueue(3)
			newQueue.EnQueue(4)

			if newQueue.IsEmptyQueue() != false {
				t.Errorf("Test Queue isEmpty is wrong the result must be %v but got %v", false, newQueue.IsEmptyQueue())
			}
		})

		t.Run("Test Queue Length", func(t *testing.T) {
			newQueue := NewQueueListArray()
			if newQueue.LenQueue() != 0 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 0, newQueue.LenQueue())
			}

			newQueue.EnQueue(3)
			newQueue.EnQueue(4)
			newQueue.DeQueue()
			newQueue.EnQueue(22)
			newQueue.EnQueue(99)
			newQueue.DeQueue()
			newQueue.DeQueue()

			if newQueue.LenQueue() != 1 {
				t.Errorf("Test Queue Length is wrong the result must be %v but got %v", 1, newQueue.LenQueue())
			}

		})
	})
}
