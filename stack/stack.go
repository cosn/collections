// Package stack implements a stack.
package stack

// S is the internal representation of the data structure.
type S struct {
	storage []interface{}
	i       int
}

// Init initializes the stack data structure.
// A stack must be initialized before it can be used.
// O(1)
func (s *S) Init(size int) {
	// the stack needs to have at least capacity for 1 element
	// defaulting here instead of panicing to allow for benchmarking
	if size == 0 {
		size = 1
	}

	s.storage = make([]interface{}, size)
	s.i = -1
}

// Push adds a new element to the top of the stack.
// O(1)
func (s *S) Push(v interface{}) {
	// dynamically increase the size of storage as needed
	if s.i+1 == cap(s.storage) {
		ns := make([]interface{}, cap(s.storage)*2)
		copy(ns, s.storage)
		s.storage = ns
	}

	s.i++
	s.storage[s.i] = v
}

// Pop removes the top element from the stack.
// O(1).
func (s *S) Pop() interface{} {
	if s.i < 0 {
		return nil
	}

	v := s.storage[s.i]
	s.storage[s.i] = nil
	s.i--

	return v
}

// Peek returns the top element from the stack without removing it.
// O(1)
func (s *S) Peek() interface{} {
	if s.i < 0 {
		return nil
	}

	return s.storage[s.i]
}

// IsEmpty returns true if the stack has no elements.
// O(1)
func (s *S) IsEmpty() bool {
	return s.Len() == 0
}

// Len returns the number of elements in the stack.
// O(1)
func (s *S) Len() int {
	return s.i + 1
}
