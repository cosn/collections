// Package tst implements a ternary search tree.
package tst

// TST is the internal representation of a ternary search tree.
type TST struct {
	r *node
	w int
}

// node is the internal represnetation of a ternary saerch tree node.
type node struct {
	lo, eq, hi, p *node
	c             rune
	e             bool
}

// Inserts adds a new word to the tree.
// O(len(s))
func (t *TST) Insert(s string) {
	t.r = insert(t.r, nil, s, t)
}

// insert recusively adds a word to the tree.
func insert(n, p *node, s string, t *TST) *node {
	if len(s) == 0 {
		if n != nil && n.p != nil && !n.p.e {
			n.p.e = true
			t.w++
		}

		return n
	}

	c := rune(s[0])
	if n == nil {
		n = &node{c: c, p: p}
		if len(s) == 1 {
			n.e = true
			t.w++
		}
	}

	if c < n.c {
		n.lo = insert(n.lo, n, s, t)
	} else if c > n.c {
		n.hi = insert(n.hi, n, s, t)
	} else {
		n.eq = insert(n.eq, n, s[1:len(s)], t)
	}

	return n
}

func (t *TST) Delete(s string) bool {
	f, n := traverse(t.r, s)

	if !f {
		return false
	}

	for n != nil {
		if !n.hasChildren() {
			if n.p == nil {
				// the node is the root, so just remove it
				n = nil
			} else if n.p != nil {
				// remove the link from the parent
				if n.p.eq == n {
					n.p.eq = nil
				} else if n.p.lo == n {
					n.p.lo = nil
				} else if n.p.hi == n {
					n.p.hi = nil
				}

				// if the parent isn't a terminating node, move up
				// otherwise, stop where we are
				if !n.p.e {
					n = n.p
				} else {
					n = nil
				}
			}
		} else {
			// the node has children, so just mark it as non-terminating
			n.e = false
			break
		}
	}

	t.w--
	return true
}

// Has returns true if the tree contains the given word.
// O(len(s))
func (t *TST) Has(s string) bool {
	f, _ := traverse(t.r, s)

	return f
}

// StartsWith returns all the words in the trie that begin with
// the given string.
// O(N)
func (t *TST) StartsWith(s string) (matches []string) {
	f, n := traverse(t.r, s)
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

	ns := s + string(n.c)

	if n.e {
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
		if rune(s[i]) < n.c {
			n = n.lo
		} else if rune(s[i]) > n.c {
			n = n.hi
		} else {
			if i++; i == len(s) {
				return n.e, n
			}
			n = n.eq
		}
	}

	return false, n
}

// Clear removes all the elements from the tree.
// O(1)
func (t *TST) Clear() {
	t.r = nil
	t.w = 0
}

// Len returns the number of words in the tree.
// O(1)
func (t *TST) Len() int {
	return t.w
}

// hasChildren returns true if the node has any children
func (n *node) hasChildren() bool {
	return n.lo != nil || n.eq != nil || n.hi != nil
}
