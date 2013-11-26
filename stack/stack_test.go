package stack

import "testing"

const iterations = 1024

func TestPushPop(t *testing.T) {
	s := new(Stack)
	s.Init(iterations / 2)

	for i := 0; i < iterations; i++ {
		s.Push(i)
	}

	for i := iterations - 1; i <= 0; i-- {
		testPop(t, s, i)
	}
}

func TestPeek(t *testing.T) {
	s := new(Stack)
	s.Init(10)

	s.Push("a")
	testPeek(t, s, "a")

	s.Push("b")
	testPeek(t, s, "b")

	s.Pop()
	testPeek(t, s, "a")

	s.Pop()
	testPeek(t, s, nil)
}

func TestLen(t *testing.T) {
	s := new(Stack)
	s.Init(iterations / 4)

	for i := 0; i < iterations; i++ {
		s.Push(i)
	}

	if l := s.Len(); l != iterations {
		t.Errorf("Stack length was expected to be %v, but is %v", iterations, l)
	}

	s.Pop()
	if l := s.Len(); l != iterations-1 {
		t.Errorf("Stack length was expected to be %v, but is %v", iterations-1, l)
	}
}

func TestIsEmpty(t *testing.T) {
	s := new(Stack)
	s.Init(2)

	if s.IsEmpty() != true {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)

	if s.IsEmpty() != false {
		t.Errorf("Stack should not be empty")
	}
}

func testPop(t *testing.T, s *Stack, e interface{}) {
	if v := s.Pop(); v != e {
		t.Errorf("Popping expected %v, got %v", e, v)
	}
}

func testPeek(t *testing.T, s *Stack, e interface{}) {
	if v := s.Peek(); v != e {
		t.Errorf("Peeking expected %v, got %v", e, v)
	}
}
