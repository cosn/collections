// Package set implements a hashset.
package set

// Set is the internal representation of the data structure.
type Set struct {
	s map[interface{}]bool
}

// Init initializes the set data structure.
// A set must be initialized before it can be used.
// O(1)
func (s *Set) Init() {
	s.s = make(map[interface{}]bool)
}

// Union returns the union of sets s and t in a new set.
// O(n+m)
func (s *Set) Union(t *Set) *Set {
	if t == nil {
		return s
	}

	ns := new(Set)
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
// O(min(n,m)
func (s *Set) Intersect(t *Set) *Set {
	if t == nil {
		return nil
	}

	ns := new(Set)
	ns.Init()

	// find the smaller Set to iterate through
	var ss, ls *Set
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
func (s *Set) Diff(t *Set) *Set {
	if t == nil {
		return s
	}

	ns := new(Set)
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
func (s *Set) SymetricDiff(t *Set) *Set {
	ns := s.Diff(t)
	nt := t.Diff(s)

	return ns.Union(nt)
}

// IsSubset returns true if set s is a subset of set t.
// O(n)
func (s *Set) IsSubset(t *Set) bool {
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
func (s *Set) IsProperSubset(t *Set) bool {
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
func (s *Set) Equals(t *Set) bool {
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

// Contains returns true if the set contains the given element.
// O(1)
func (s *Set) Contains(e interface{}) bool {
	return s.s[e]
}

// Len returns the number of elements in the set.
// O(1)
func (s *Set) Len() int {
	return len(s.s)
}

// IsEmpty returns true if the set contains any elements.
// O(1)
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

// Clear removes all elements from the set.
// O(n)
func (s *Set) Clear() {
	for v := range s.s {
		delete(s.s, v)
	}
}

// Iter provides an iterator over the set.
// O(n)
func (s *Set) Iter() <-chan interface{} {
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
func (s *Set) Add(e interface{}) (f bool) {
	if !s.s[e] {
		s.s[e] = true
		f = true
	}

	return
}

// Remove removes an element from the set and returns true if the value previously existed.
// O(1)
func (s *Set) Remove(e interface{}) (f bool) {
	_, f = s.s[e]
	delete(s.s, e)

	return
}
