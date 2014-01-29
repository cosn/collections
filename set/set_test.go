package set

import "testing"

func TestAddRemoveHasClear(t *testing.T) {
	s := new(S)
	s.Init()

	testLen(t, s.Len(), 0)
	if !s.IsEmpty() {
		t.Errorf("Set should be empty, but it doesn't seem to be\n")
	}

	s.Add("a")
	testLen(t, s.Len(), 1)
	if !s.Has("a") {
		t.Errorf("'a' was not found in the set\n")
	}

	// duplicate value should not be added
	if s.Add("a") {
		t.Errorf("'a' already exited in the set, so it should not be added again\n")
	}
	testLen(t, s.Len(), 1)

	s.Add("b")
	testLen(t, s.Len(), 2)
	if !s.Has("a") {
		t.Errorf("'a' was not found in the set\n")
	}
	if !s.Has("b") {
		t.Errorf("'b' was not found in the set\n")
	}

	s.Add("c")
	s.Remove("b")
	testLen(t, s.Len(), 2)
	if !s.Has("a") {
		t.Errorf("'a' was not found in the set\n")
	}
	if s.Has("b") {
		t.Errorf("'b' was found in the set, but should not have been\n")
	}
	if !s.Has("c") {
		t.Errorf("'c' was not found in the set\n")
	}

	if s.Remove("b") {
		t.Errorf("'b' doesn't exist, so it should not be removed again\n")
	}

	s.Clear()
	testLen(t, s.Len(), 0)
}

func TestUnion(t *testing.T) {
	s1, s2 := new(S), new(S)
	s1.Init()
	s2.Init()

	s1.Add(0)
	s1.Add(1)

	s2.Add(2)
	s2.Add(3)

	s := s1.Union(s2)

	testLen(t, s.Len(), 4)

	for i := 0; i <= 3; i++ {
		if !s.Has(i) {
			t.Errorf("Element '%v' not found\n", i)
		}
	}

	ns := s1.Union(nil)
	if ns != s1 {
		t.Errorf("Union with null set should be the set\n")
	}
}

func TestIntersect(t *testing.T) {
	s1, s2 := new(S), new(S)
	s1.Init()
	s2.Init()

	s1.Add(0)
	s1.Add(1)
	s1.Add(2)
	s1.Add(5)

	s2.Add(1)
	s2.Add(2)
	s2.Add(6)

	s := s1.Intersect(s2)

	testLen(t, s.Len(), 2)

	for i := 1; i <= 2; i++ {
		if !s.Has(i) {
			t.Errorf("Element '%v' not found\n", i)
		}
	}

	ns := s1.Intersect(nil)
	if !ns.IsEmpty() {
		t.Errorf("Intersection with null set should be an empty set\n")
	}

	e := new(S)
	es := s1.Intersect(e)
	if !es.IsEmpty() {
		t.Errorf("Intersection with empty set should be an empty set\n")
	}
}

func TestDiff(t *testing.T) {
	s1, s2 := new(S), new(S)
	s1.Init()
	s2.Init()

	s1.Add(0)
	s1.Add(1)
	s1.Add(2)

	s2.Add(1)
	s2.Add(2)
	s2.Add(3)

	s := s1.Diff(s2)

	testLen(t, s.Len(), 1)

	if !s.Has(0) {
		t.Errorf("Element '%v' not found\n", 0)
	}

	ns := s1.Diff(nil)
	if ns != s1 {
		t.Errorf("Difference with null set should be a the set\n")
	}
}

func TestSymetricDiff(t *testing.T) {
	s1, s2 := new(S), new(S)
	s1.Init()
	s2.Init()

	s1.Add(0)
	s1.Add(2)
	s1.Add(3)

	s2.Add(1)
	s2.Add(2)
	s2.Add(3)
	s2.Add(5)

	s := s1.SymetricDiff(s2)

	testLen(t, s.Len(), 3)

	if !s.Has(0) {
		t.Errorf("Element '%v' not found\n", 0)
	}

	if !s.Has(1) {
		t.Errorf("Element '%v' not found\n", 1)
	}

	if !s.Has(5) {
		t.Errorf("Element '%v' not found\n", 5)
	}
}

func TestIsSubset(t *testing.T) {
	s1, s2 := new(S), new(S)
	s1.Init()
	s2.Init()

	s1.Add("a")
	s1.Add("b")
	s1.Add("c")

	s2.Add("a")
	s2.Add("b")
	s2.Add("c")

	if !s1.IsSubset(s2) {
		t.Errorf("S1 is not a subset of S2 as expected\n")
	}

	if s1.IsSubset(nil) {
		t.Errorf("The set should not be a subset of an empty set\n")
	}
}

func TestIsProperSubset(t *testing.T) {
	s1, s2 := new(S), new(S)
	s1.Init()
	s2.Init()

	s1.Add("a")
	s1.Add("c")

	s2.Add("a")
	s2.Add("b")
	s2.Add("c")

	if !s1.IsProperSubset(s2) {
		t.Errorf("S1 is not a subset of S2 as expected\n")
	}

	if s1.IsProperSubset(nil) {
		t.Errorf("The set should not be a proper subset of an empty set\n")
	}
}

func TestEquals(t *testing.T) {
	s1, s2 := new(S), new(S)
	s1.Init()
	s2.Init()

	s1.Add("a")
	s1.Add("b")
	s1.Add("c")

	s2.Add("a")
	s2.Add("b")
	s2.Add("c")

	if !s1.Equals(s2) {
		t.Errorf("S1 is not equal to S2\n")
	}

	s2.Add("d")
	if s1.Equals(s2) {
		t.Errorf("S1 should not be equal to S2\n")
	}

	s1.Add("d")
	s1.Add("e")
	if s1.Equals(s2) {
		t.Errorf("S1 should not be equal to S2\n")
	}

	if s1.Equals(nil) {
		t.Errorf("The set should not equal the empty set\n")
	}
}

func TestIter(t *testing.T) {
	s := new(S)
	s.Init()

	m := make(map[interface{}]bool)
	for i := 0; i < 100; i++ {
		s.Add(i)
		m[i] = false
	}

	for e := range s.Iter() {
		m[e] = true
	}

	for k, v := range m {
		if !v {
			t.Errorf("Element %v was not interated in the set\n", k)
		}
	}
}

func testLen(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("Expected set length to be %v, instead was %v\n", expected, actual)
	}
}
