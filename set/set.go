// Package set implements a hashset.
package set

// S is the internal representation of the data structure.
type S struct {
	s map[interface{}]bool
}

// Init initializes the set data structure.
// A set must be initialized before it can be used.
// O(1)
func (s *S) Init() {
	s.s = make(map[interface{}]bool)
}

// Union returns the union of sets s and t in a new set.
// O(n+m)
func (s *S) Union(t *S) *S {
	if t == nil {
		return s
	}

	ns := new(S)
	ns.Init()

	for e := range s.s {
		ns.Add(e)
	}

	for e := range t.s {
		ns.Add(e)
	}

	return ns
}

// Intersect returns the intersection of sets s and t in a new set.
// If t is null or empty, the result is an empty set
// O(min(n,m)
func (s *S) Intersect(t *S) *S {
	ns := new(S)
	ns.Init()

	if t == nil || t.Len() == 0 {
		return ns
	}

	// find the smaller Set to iterate through
	var ss, ls *S
	if s.Len() < t.Len() {
		ss = s
		ls = t
	} else {
		ss = t
		ls = s
	}

	for e := range ss.s {
		if ls.s[e] {
			ns.Add(e)
		}
	}

	return ns
}

// Diff returns the difference between sets s and t in a new set.
// O(n)
func (s *S) Diff(t *S) *S {
	if t == nil || t.Len() == 0 {
		return s
	}

	ns := new(S)
	ns.Init()

	for e := range s.s {
		if !t.s[e] {
			ns.Add(e)
		}
	}

	return ns
}

// SymetricDiff returns a new set with elements from one set or the other, but not both.
// O(n+m)
func (s *S) SymetricDiff(t *S) *S {
	ns := s.Diff(t)
	nt := t.Diff(s)

	return ns.Union(nt)
}

// IsSubset returns true if set s is a subset of set t.
// O(n)
func (s *S) IsSubset(t *S) bool {
	if t == nil {
		return false
	}

	if s.Len() > t.Len() {
		return false
	}

	for e := range s.s {
		if !t.s[e] {
			return false
		}
	}

	return true
}

// IsSubset returns true if set s is a proper subset of set t.
// O(n)
func (s *S) IsProperSubset(t *S) bool {
	if t == nil {
		return false
	}

	if s.Len() >= t.Len() {
		return false
	}

	for e := range s.s {
		if !t.s[e] {
			return false
		}
	}

	return true
}

// Equal returns true if the two sets contain the same elements
// O(n)
func (s *S) Equals(t *S) bool {
	if t == nil || s.Len() != t.Len() {
		return false
	}

	for e := range s.s {
		if !t.s[e] {
			return false
		}
	}

	return true
}

// Has returns true if the set contains the given element.
// O(1)
func (s *S) Has(e interface{}) bool {
	return s.s[e]
}

// Len returns the number of elements in the set.
// O(1)
func (s *S) Len() int {
	return len(s.s)
}

// IsEmpty returns true if the set contains any elements.
// O(1)
func (s *S) IsEmpty() bool {
	return s.Len() == 0
}

// Clear removes all elements from the set.
// O(n)
func (s *S) Clear() {
	for v := range s.s {
		delete(s.s, v)
	}
}

// Iter provides an iterator over the set.
// O(n)
func (s *S) Iter() <-chan interface{} {
	c := make(chan interface{}, s.Len())
	go func() {
		for e := range s.s {
			c <- e
		}
		close(c)
	}()

	return c
}

// Add adds a new element to the set and returns true if the value previously existed.
// O(1)
func (s *S) Add(e interface{}) (f bool) {
	if !s.s[e] {
		s.s[e] = true
		f = true
	}

	return
}

// Remove removes an element from the set and returns true if the value previously existed.
// O(1)
func (s *S) Remove(e interface{}) (f bool) {
	_, f = s.s[e]
	delete(s.s, e)

	return
}
