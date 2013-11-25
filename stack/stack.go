// Package stack implements a stack
package stack

// stack is the internal representation of the data structure
type stack struct {
	storage []interface{}
	i       int
}

// New returns an initialized stack of a given size
func New(size int) *stack {
	s := new(stack)
	s.storage = make([]interface{}, size)
	s.i = -1

	return s
}

// Push adds a new element to the top of the stack
// O(1)
func (s *stack) Push(v interface{}) {
	// dynamically increase the size of storage as needed
	if s.i+1 == cap(s.storage) {
		ns := make([]interface{}, s.i*2)
		copy(ns, s.storage)
		s.storage = ns
	}

	s.i++
	s.storage[s.i] = v
}

// Pop removes the top element from the stack
// O(1).
func (s *stack) Pop() interface{} {
	if s.i < 0 {
		return nil
	}

	v := s.storage[s.i]
	s.i--

	return v
}

// Peek returns the top element from the stack without removing it
// O(1)
func (s *stack) Peek() interface{} {
	if s.i < 0 {
		return nil
	}

	return s.storage[s.i]
}

// IsEmpty returns a value indicating whether the stack has any elements
// O(1)
func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

// Len returns the number of elements in the stack
// O(1)
func (s *stack) Len() int {
	return s.i + 1
}
