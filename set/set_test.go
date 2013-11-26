package set

import "testing"

func TestAddRemoveContainsClear(t *testing.T) {
	s := new(Set)
	s.Init()

	testLen(t, s.Len(), 0)
	if !s.IsEmpty() {
		t.Errorf("Set should be empty, but it doesn't seem to be")
	}

	s.Add("a")
	testLen(t, s.Len(), 1)
	if !s.Contains("a") {
		t.Errorf("'a' was not found in the set\n")
	}

	// duplicate value should not be added
	if s.Add("a") {
		t.Errorf("'a' already exited in the set, so it should not be added again")
	}
	testLen(t, s.Len(), 1)

	s.Add("b")
	testLen(t, s.Len(), 2)
	if !s.Contains("a") {
		t.Errorf("'a' was not found in the set\n")
	}
	if !s.Contains("b") {
		t.Errorf("'b' was not found in the set\n")
	}

	s.Add("c")
	s.Remove("b")
	testLen(t, s.Len(), 2)
	if !s.Contains("a") {
		t.Errorf("'a' was not found in the set\n")
	}
	if s.Contains("b") {
		t.Errorf("'b' was found in the set, but should not have been\n")
	}
	if !s.Contains("c") {
		t.Errorf("'c' was not found in the set\n")
	}

	if s.Remove("b") {
		t.Errorf("'b' doesn't exist, so it should not be removed again")
	}

	s.Clear()
	testLen(t, s.Len(), 0)
}

func TestUnion(t *testing.T) {
	s1, s2 := new(Set), new(Set)
	s1.Init()
	s2.Init()

	s1.Add(0)
	s1.Add(1)

	s2.Add(2)
	s2.Add(3)

	s := s1.Union(s2)

	testLen(t, s.Len(), 4)

	for i := 0; i <= 3; i++ {
		if !s.Contains(i) {
			t.Errorf("Element '%v' not found\n", i)
		}
	}

	ns := s1.Union(nil)
	if ns != s1 {
		t.Errorf("Union with null set should be the set")
	}
}

func TestIntersect(t *testing.T) {
	s1, s2 := new(Set), new(Set)
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
		if !s.Contains(i) {
			t.Errorf("Element '%v' not found\n", i)
		}
	}

	ns := s1.Intersect(nil)
	if ns != nil {
		t.Errorf("Intersection with null set should be a null set")
	}
}

func TestDiff(t *testing.T) {
	s1, s2 := new(Set), new(Set)
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

	if !s.Contains(0) {
		t.Errorf("Element '%v' not found\n", 0)
	}

	ns := s1.Diff(nil)
	if ns != s1 {
		t.Errorf("Difference with null set should be a the set")
	}
}

func TestSymetricDiff(t *testing.T) {
	s1, s2 := new(Set), new(Set)
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

	if !s.Contains(0) {
		t.Errorf("Element '%v' not found\n", 0)
	}

	if !s.Contains(1) {
		t.Errorf("Element '%v' not found\n", 1)
	}

	if !s.Contains(5) {
		t.Errorf("Element '%v' not found\n", 5)
	}
}

func TestIsSubset(t *testing.T) {
	s1, s2 := new(Set), new(Set)
	s1.Init()
	s2.Init()

	s1.Add("a")
	s1.Add("b")
	s1.Add("c")

	s2.Add("a")
	s2.Add("b")
	s2.Add("c")

	if !s1.IsSubset(s2) {
		t.Errorf("S1 is not a subset of S2 as expected")
	}

	if s1.IsSubset(nil) {
		t.Errorf("The set should not be a subset of an empty set")
	}
}

func TestIsProperSubset(t *testing.T) {
	s1, s2 := new(Set), new(Set)
	s1.Init()
	s2.Init()

	s1.Add("a")
	s1.Add("c")

	s2.Add("a")
	s2.Add("b")
	s2.Add("c")

	if !s1.IsProperSubset(s2) {
		t.Errorf("S1 is not a subset of S2 as expected")
	}

	if s1.IsProperSubset(nil) {
		t.Errorf("The set should not be a proper subset of an empty set")
	}
}

func TestEqual(t *testing.T) {
	s1, s2 := new(Set), new(Set)
	s1.Init()
	s2.Init()

	s1.Add("a")
	s1.Add("b")
	s1.Add("c")

	s2.Add("a")
	s2.Add("b")
	s2.Add("c")

	if !s1.Equal(s2) {
		t.Errorf("S1 is not equal to S2")
	}

	s2.Add("d")
	if s1.Equal(s2) {
		t.Errorf("S1 should not be equal to S2")
	}

	s1.Add("d")
	s1.Add("e")
	if s1.Equal(s2) {
		t.Errorf("S1 should not be equal to S2")
	}

	if s1.Equal(nil) {
		t.Errorf("The set should not equal the empty set")
	}
}

func TestIter(t *testing.T) {
	s := new(Set)
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
			t.Errorf("Element %v was not interated in the set", k)
		}
	}
}

func testLen(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("Expected set length to be %v, instead was %v\n", expected, actual)
	}
}
