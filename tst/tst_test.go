package tst

import (
	"fmt"
	"testing"
)

func TestInsertHas(t *testing.T) {
	values := []string{"hey", "hello", "hell", "magazine", "magnificent", "magazines"}
	tst := new(TST)

	for _, s := range values {
		tst.Insert(s)
	}

	for _, s := range values {
		if !tst.Has(s) {
			t.Errorf("Tree expected to contain '%v', but did not", s)
		}
	}

	if tst.Has("mag") {
		t.Error("Unexpected value found in tree")
	}
}

func TestStartsWith(t *testing.T) {
	tst := new(TST)

	tst.Insert("foo")
	tst.Insert("foobar")
	tst.Insert("f")
	tst.Insert("fo")

	if m := tst.StartsWith("foo"); len(m) != 2 {
		t.Errorf("Did not find the expected two matches: %v", m)
	}

	if m := tst.StartsWith("fo"); len(m) != 3 {
		t.Errorf("Did not find the expected three matches: %v", m)
	}

	if m := tst.StartsWith("f"); len(m) != 4 {
		t.Errorf("Did not find the expected four matches: %v", m)
	}
}

func TestLenClear(t *testing.T) {
	tst := new(TST)

	tst.Insert("hello")
	tst.Insert("hell")
	tst.Insert("hey")
	tst.Insert("heck")
	tst.Insert("blah")
	tst.Insert("boo")
	tst.Insert("foo")
	tst.Insert("foobar")
	tst.Insert("moo")

	if l := tst.Len(); l != 9 {
		t.Errorf("Tree length expected to be 9, but instead was %v", l)
	}

	tst.Clear()
	if l := tst.Len(); l != 0 {
		t.Errorf("Tree expected to be empty, instead has %v elements", l)
	}
}

func TestDelete(t *testing.T) {
	values := []string{"hey", "hello", "hell", "magazine", "magnificent", "magazines"}
	expected := []string{"hey", "hello", "magazine", "magnificent"}

	tst := new(TST)

	for _, s := range values {
		tst.Insert(s)
	}

	if !tst.Delete("magazines") {
		t.Error("Value should have been removed")
	}

	if tst.Has("magazines") {
		t.Error("Word should have been removed, but is still in tree")
	}

	if !tst.Delete("hell") {
		t.Error("Value should have been removed")
	}

	if tst.Has("hell") {
		t.Error("Word should have been removed, but is still in tree")
	}

	if l := tst.Len(); l != len(expected) {
		t.Errorf("Number of words should be %v, but instead was %v", len(expected), l)
	}

	for _, s := range expected {
		if !tst.Has(s) {
			t.Errorf("Tree expected to contain '%v', but did not", s)
		}
	}
}

func (t *TST) String() (s string) {
	s = fmt.Sprintf("%v\n", t.words)
	print(t.root, "", "root: ", &s)
	return
}

func print(n *node, space, pos string, s *string) {
	if n == nil {
		return
	}

	*s += fmt.Sprintf("%v%q %t (%p)\n", space+pos, n.char, n.end, &n)

	space += " "
	print(n.lo, space, "lo: ", s)
	print(n.eq, space, "eq: ", s)
	print(n.hi, space, "hi: ", s)
}
