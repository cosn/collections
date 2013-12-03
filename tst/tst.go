// Package tst implements a ternary search tree.
package tst

// TST is the internal representation of a ternary search tree.
type TST struct {
	root  *node
	words int
}

// node is the internal represnetation of a ternary saerch tree node.
type node struct {
	lo, eq, hi, parent *node
	char               rune
	end                bool
}

// Inserts adds a new word to the tree.
// Average: O(log(n)) Worst: O(n)
func (t *TST) Insert(s string) {
	t.root = insert(t.root, nil, s, t)
}

// insert recusively adds a word to the tree.
func insert(n, p *node, s string, t *TST) *node {
	if len(s) == 0 {
		if n != nil && n.parent != nil && !n.parent.end {
			n.parent.end = true
			t.words++
		}

		return n
	}

	c := rune(s[0])
	if n == nil {
		n = &node{char: c, parent: p}
		if len(s) == 1 {
			n.end = true
			t.words++
		}
	}

	if c < n.char {
		n.lo = insert(n.lo, n, s, t)
	} else if c > n.char {
		n.hi = insert(n.hi, n, s, t)
	} else {
		n.eq = insert(n.eq, n, s[1:len(s)], t)
	}

	return n
}

// Delete returns true if the word was removed from the tree.
// Average: O(log(n)) Worst: O(n)
func (t *TST) Delete(s string) bool {
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
func (t *TST) Has(s string) bool {
	f, _ := traverse(t.root, s)

	return f
}

// StartsWith returns all the words in the trie that begin with
// the given string.
// O(n)
func (t *TST) StartsWith(s string) (matches []string) {
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
func (t *TST) Clear() {
	t.root = nil
	t.words = 0
}

// Len returns the number of words in the tree.
// O(1)
func (t *TST) Len() int {
	return t.words
}

// hasChildren returns true if the node has any children
func (n *node) hasChildren() bool {
	return n.lo != nil || n.eq != nil || n.hi != nil
}
