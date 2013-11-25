package queue

import "testing"

const iterations = 1000

func TestPushPop(t *testing.T) {
	q := New()

	for i := 0; i < iterations; i++ {
		q.Push(i)
	}

	for i := 0; i < iterations; i++ {
		testPop(t, q, i)
	}
}

func TestLen(t *testing.T) {
	q := New()

	for i := 0; i < iterations; i++ {
		q.Push(i)
	}

	if l := q.Len(); l != iterations {
		t.Errorf("Queue length was expected to be %v, but is %v", iterations, l)
	}

	q.Pop()
	if l := q.Len(); l != iterations-1 {
		t.Errorf("Queue length was expected to be %v, but is %v", iterations-1, l)
	}
}

func TestIsEmpty(t *testing.T) {
	q := New()

	if q.IsEmpty() != true {
		t.Errorf("Queue should be empty")
	}

	q.Push(1)

	if q.IsEmpty() != false {
		t.Errorf("Queue should not be empty")
	}
}

func testPop(t *testing.T, q *queue, e interface{}) {
	if v := q.Pop(); v != e {
		t.Errorf("Popping expected %v, got %v", e, v)
	}
}
