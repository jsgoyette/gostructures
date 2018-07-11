package queue

import (
	"fmt"

	"github.com/altay13/godatastructures/doublylinkedlist"
)

// Queue implementation
type Queue struct {
	list *doublylinkedlist.List
	size int

	maxSize int
}

// NewQueue method creates the stack
func NewQueue() *Queue {
	l := doublylinkedlist.NewList()
	q := &Queue{
		list:    l,
		maxSize: -1,
	}
	return q
}

func (q *Queue) SetSize(size int) error {
	if q.size != 0 {
		return fmt.Errorf("[ ERR ] Queue is not empty. To set the size, queue should be empty.")
	}

	q.maxSize = size
	return nil
}

// Enqueue ...
func (q *Queue) Enqueue(item interface{}) {
	if q.size == q.maxSize {
		q.Dequeue()
	}
	q.list.AddNode(item)
	q.size++
}

// Dequeue ...
func (q *Queue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	}
	n := q.list.GetFirstNode()
	q.list.RemoveFirstNode()
	if n != nil {
		q.size--
	}
	return n.Item
}
