// Package queue implements a queue
package queue

import "container/list"

// queue is the internal representation of the data structure
type queue struct {
	l *list.List
}

// New returns an initialized queue
func New() *queue {
	q := new(queue)
	q.l = list.New()
	return q
}

// Push enqueues an element to the queue
// O(1)
func (s *queue) Push(v interface{}) {
	s.l.PushFront(v)
}

// Pop dequeues an element from the queue
// O(1)
func (s *queue) Pop() interface{} {
	if s.l.Len() == 0 {
		return nil
	}

	v := s.l.Back()
	return s.l.Remove(v)
}

// Len returns the number of elements in the queue
func (s *queue) Len() int {
	return s.l.Len()
}

// IsEmpty returns a value indicating whether the queue has any elements
func (s *queue) IsEmpty() bool {
	return s.l.Len() == 0
}
