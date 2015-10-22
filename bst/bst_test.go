package bst

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	expected := []int{5, 3, 7, 4, 6}
	bst := new(T)

	for _, i := range expected {
		if !bst.Insert(i, i) {
			t.Errorf("Element %v should have been added to the tree", i)
		}
	}

	for _, i := range expected {
		if bst.Find(i) == nil {
			t.Errorf("Element %v expected to be in the tree, but was not", i)
		}
	}

	if bst.Insert(4, 44) {
		t.Error("Duplicate elements should not be added")
	}

	if bst.Find(4) == 44 {
		t.Error("Previously inserted elements should not be updated")
	}

	if c := bst.count; c != len(expected) {
		t.Errorf("Tree expected to have %v elements, but has %v instead", len(expected), c)
	}
}

func TestRemove_SingleElement(t *testing.T) {
	bst := new(T)

	bst.Insert(5, 10)

	if !bst.Delete(5) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(5) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}
}

func TestRemove_RootWithSingleChild(t *testing.T) {
	bst := new(T)

	bst.Insert(5, 10)
	bst.Insert(4, 8)

	if !bst.Delete(5) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(5) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}

	if bst.Find(4) == nil {
		t.Errorf("Element with key %v was not found", 4)
	}

	if bst.count != 1 {
		t.Errorf("Expected element count %v found cound %v", 1, bst.count)
	}
}

func TestRemove_RootWithTwoChildren(t *testing.T) {
	bst := new(T)

	bst.Insert(5, 10)
	bst.Insert(4, 8)
	bst.Insert(6, 12)

	if !bst.Delete(5) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(5) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}

	if bst.Find(4) == nil {
		t.Errorf("Element with key %v was not found", 4)
	}

	if bst.Find(6) == nil {
		t.Errorf("Element with key %v was not found", 6)
	}

	if bst.count != 2 {
		t.Errorf("Expected element count %v found cound %v", 1, bst.count)
	}
}

func TestRemove(t *testing.T) {
	expected := []int{5, 3, 7, 4, 6}
	bst := new(T)

	for _, i := range expected {
		bst.Insert(i, i)
	}

	if !bst.Delete(6) {
		t.Errorf("Element %v should have been removed from the tree", 6)
	}

	for _, i := range expected[0:3] {
		if bst.Find(i) == nil {
			t.Errorf("Element %v expected to be in the tree, but was not", i)
		}
	}

	if d := expected[len(expected)-1]; bst.Find(d) != nil {
		t.Errorf("Element %v should have been removed from the tree", d)
	}

	if bst.Delete(6) {
		t.Error("Duplicate elements should not be delete")
	}

	if c := bst.count; c != len(expected)-1 {
		t.Errorf("Tree expected to have %v elements, but has %v instead", len(expected)-1, c)
	}
}

func TestTraverse_InOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{3, 4, 5, 6, 7}

	bst := new(T)

	for _, i := range elements {
		bst.Insert(i, i)
	}

	i := 0
	for e := range bst.Traverse(InOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestTraverse_PreOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{5, 3, 4, 7, 6}

	bst := new(T)

	for _, i := range elements {
		bst.Insert(i, i)
	}

	i := 0
	for e := range bst.Traverse(PreOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestTraverse_PostOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{4, 3, 6, 7, 5}

	bst := new(T)

	for _, i := range elements {
		bst.Insert(i, i)
	}

	i := 0
	for e := range bst.Traverse(PostOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestClear(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}

	bst := new(T)

	for _, i := range elements {
		bst.Insert(i, i)
	}

	bst.Clear()

	if c := bst.count; c != 0 {
		t.Errorf("Expected tree to be empty, but has %v elements", c)
	}

	if p := bst.String(); len(p) != 0 {
		t.Errorf("No elements expected in the tree, but found %v", p)
	}
}

func (t *T) String() (s string) {
	print(t.root, &s)
	return
}

func print(n *node, s *string) {
	if n == nil {
		return
	}

	*s += fmt.Sprintf("%p %v\n", n, n)
	print(n.l, s)
	print(n.r, s)
}

func BenchmarkInsert(b *testing.B) {
	bst := new(T)
	for _, i := range rand.Perm(b.N) {
		bst.Insert(i, i)
	}
}

func BenchmarkDelete(b *testing.B) {
	bst := new(T)
	for _, i := range rand.Perm(b.N) {
		bst.Insert(i, i)
	}

	b.ResetTimer()
	for _, i := range rand.Perm(b.N) {
		bst.Delete(i)
	}
}

func BenchmarkFind(b *testing.B) {
	bst := new(T)
	for _, i := range rand.Perm(b.N) {
		bst.Insert(i, i)
	}

	b.ResetTimer()
	for _, i := range rand.Perm(b.N) {
		bst.Find(i)
	}
}
