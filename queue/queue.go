// Package queue implements a queue.
package queue

import "container/list"

// Q is the internal representation of the data structure.
type Q struct {
	l *list.List
}

// Init initializes the queue data structure.
// A queue must be initialized before it can be used.
// O(1)
func (q *Q) Init() {
	q.l = list.New()
}

// Push enqueues an element to the queue.
// O(1)
func (q *Q) Push(v interface{}) {
	q.l.PushFront(v)
}

// Pop dequeues an element from the queue.
// O(1)
func (q *Q) Pop() interface{} {
	if q.l.Len() == 0 {
		return nil
	}

	v := q.l.Back()
	return q.l.Remove(v)
}

// Len returns the number of elements in the queue.
// O(1)
func (q *Q) Len() int {
	return q.l.Len()
}

// IsEmpty returns true the queue has no elements.
// O(1)
func (q *Q) IsEmpty() bool {
	return q.l.Len() == 0
}
