package trie

import "testing"
import "fmt"

func TestInsertHas(t *testing.T) {
	values := []string{"hey", "hello", "hell", "magazine", "magnificent", "magazines"}
	trie := new(Trie)
	trie.Init(256)

	for _, s := range values {
		trie.Insert(s)
	}

	for _, s := range values {
		if !trie.Has(s) {
			t.Errorf("Trie expected to contain '%v', but did not", s)
		}
	}

	if trie.Has("mag") {
		t.Error("Unexpected value found in trie")
	}
}

func TestStartsWith(t *testing.T) {
	trie := new(Trie)
	trie.Init(256)

	trie.Insert("foo")
	trie.Insert("foobar")
	trie.Insert("f")
	trie.Insert("fo")

	if m := trie.StartsWith("foo"); len(m) != 2 {
		t.Errorf("Did not find the expected two matches: %v", m)
	}

	if m := trie.StartsWith("fo"); len(m) != 3 {
		t.Errorf("Did not find the expected three matches: %v", m)
	}

	if m := trie.StartsWith("f"); len(m) != 4 {
		t.Errorf("Did not find the expected four matches: %v", m)
	}
}

func TestLenClear(t *testing.T) {
	trie := new(Trie)
	trie.Init(26)

	trie.Insert("hello")
	trie.Insert("hell")
	trie.Insert("hey")
	trie.Insert("heck")
	trie.Insert("blah")
	trie.Insert("boo")
	trie.Insert("foo")
	trie.Insert("foobar")
	trie.Insert("moo")

	if l := trie.Len(); l != 9 {
		t.Errorf("Trie length expected to be 9, but instead was %v", l)
	}

	trie.Clear()
	if l := trie.Len(); l != 0 {
		t.Errorf("Trie expected to be empty, instead has %v elements", l)
	}
}

func (t *Trie) String() (s string) {
	print(t.r, "", &s)
	return
}

func print(n *node, w string, s *string) {
	if n == nil {
		return
	}

	*s += fmt.Sprintf("%v%q (%t)\n", w, n.c, n.e)

	w += " "
	for c := 0; c < len(n.n); c++ {
		print(n.n[c], w, s)
	}
}
