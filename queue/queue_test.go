package queue

import "testing"

const iterations = 1024

func TestPushPop(t *testing.T) {
	q := new(Q)
	q.Init()

	for i := 0; i < iterations; i++ {
		q.Push(i)
	}

	for i := 0; i < iterations; i++ {
		testPop(t, q, i)
	}
}

func TestLen(t *testing.T) {
	q := new(Q)
	q.Init()

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
	q := new(Q)
	q.Init()

	if q.IsEmpty() != true {
		t.Errorf("Queue should be empty")
	}

	q.Push(1)

	if q.IsEmpty() != false {
		t.Errorf("Queue should not be empty")
	}
}

func testPop(t *testing.T, q *Q, e interface{}) {
	if v := q.Pop(); v != e {
		t.Errorf("Popping expected %v, got %v", e, v)
	}
}

func BenchmarkPush(b *testing.B) {
	q := new(Q)
	q.Init()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	q := new(Q)
	q.Init()

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
