// Package tst implements a ternary search tree.
package tst

// T is the internal representation of a ternary search tree.
type T struct {
	root  *node
	words int
}

// node is the internal represnetation of a ternary saerch tree node.
type node struct {
	lo, eq, hi, parent *node
	char               rune
	end                bool
	value              interface{}
}

// Inserts adds a new word to the tree.
// The word may be accompanied by a value.
// Average: O(log(n)) Worst: O(n)
func (t *T) Insert(s string, v interface{}) {
	t.root = insert(t.root, nil, s, t, v)
}

// insert recusively adds a word to the tree.
func insert(n, p *node, s string, t *T, v interface{}) *node {
	if len(s) == 0 {
		if n != nil && n.parent != nil && !n.parent.end {
			n.parent.end = true
			n.parent.value = v
			t.words++
		}

		return n
	}

	c := rune(s[0])
	if n == nil {
		n = &node{char: c, parent: p}
		if len(s) == 1 {
			n.end = true
			n.value = v
			t.words++
		}
	}

	if c < n.char {
		n.lo = insert(n.lo, n, s, t, v)
	} else if c > n.char {
		n.hi = insert(n.hi, n, s, t, v)
	} else {
		n.eq = insert(n.eq, n, s[1:len(s)], t, v)
	}

	return n
}

// Delete returns true if the word was removed from the tree.
// Average: O(log(n)) Worst: O(n)
func (t *T) Delete(s string) bool {
	f, n := traverse(t.root, s)

	if !f {
		return false
	}

	for n != nil {
		if !n.hasChildren() {
			if n.parent == nil {
				// the node is the root, so just remove it
				n = nil
			} else {
				// remove the link from the parent
				if n.parent.eq == n {
					n.parent.eq = nil
				} else if n.parent.lo == n {
					n.parent.lo = nil
				} else if n.parent.hi == n {
					n.parent.hi = nil
				}

				// if the parent isn't a terminating node, move up
				// otherwise, stop where we are
				if !n.parent.end {
					n = n.parent
				} else {
					n = nil
				}
			}
		} else {
			// the node has children, so just mark it as non-terminating
			n.end = false
			break
		}
	}

	t.words--
	return true
}

// Has returns true if the tree contains the given word.
// Average: O(log(n)) Worst: O(n)
func (t *T) Has(s string) bool {
	_, f := t.Get(s)

	return f
}

// Get returns the value stored with the string and
// true if the tree contains the given word.
// Average: O(log(n)) Worst: O(n)
func (t *T) Get(s string) (interface{}, bool) {
	f, n := traverse(t.root, s)

	if !f {
		return nil, false
	}

	return n.value, true
}

// StartsWith returns all the words in the trie that begin with
// the given string.
// O(n)
func (t *T) StartsWith(s string) (matches []string) {
	f, n := traverse(t.root, s)
	if f {
		matches = append(matches, s)
	}
	matches = append(matches, match(n.eq, s)...)

	return
}

// match recurisvely searches for matches for a given string.
func match(n *node, s string) (matches []string) {
	if n == nil {
		return matches
	}

	ns := s + string(n.char)

	if n.end {
		matches = append(matches, ns)
	}

	matches = append(matches, match(n.eq, ns)...)
	matches = append(matches, match(n.lo, s)...)
	matches = append(matches, match(n.hi, s)...)

	return
}

// traverse returns the last matching node for a given word.
func traverse(n *node, s string) (bool, *node) {
	i := 0
	for n != nil {
		if rune(s[i]) < n.char {
			n = n.lo
		} else if rune(s[i]) > n.char {
			n = n.hi
		} else {
			if i++; i == len(s) {
				return n.end, n
			}
			n = n.eq
		}
	}

	return false, n
}

// Clear removes all the elements from the tree.
// O(1)
func (t *T) Clear() {
	t.root = nil
	t.words = 0
}

// Len returns the number of words in the tree.
// O(1)
func (t *T) Len() int {
	return t.words
}

// hasChildren returns true if the node has any children
func (n *node) hasChildren() bool {
	return n.lo != nil || n.eq != nil || n.hi != nil
}
