// Package trie implements a trie.
package trie

// Trie is the internal representation of a trie.
type Trie struct {
	r *node
	w int
	s rune
}

// node is the internal representation of a trie node.
type node struct {
	c  rune
	n  map[rune]*node
	nc int
	e  bool
	p  *node
}

// Init initializes a trie with a given alphabet size.
// A trie must be initialized before it can be used.
// O(1)
func (t *Trie) Init(size rune) {
	if size < 1 {
		panic("Trie size must be a positive number")
	}

	t.r = &node{n: make(map[rune]*node, size)}
	t.s = size
}

// Insert adds a new word to the trie.
// O(len(s))
func (t *Trie) Insert(s string) {
	r := t.r

	for i, c := range s {
		n := r.next(c, t.s)
		end := i == len(s)-1
		if n == nil {
			n = &node{c, make(map[rune]*node, t.s), 0, end, r}
			r.n[c%t.s] = n
			// increment the number of children for the parent
			// this information is useful for deletions
			r.nc++
			if end {
				t.w++
			}
		} else if end && !n.e {
			// if the node previously existed, but wasn't a terminating string,
			// we need to now mark it as such (i.e. insert("foobar"), insert("foo"))
			// also increment the number of words in the trie in this scenario
			n.e = true
			t.w++
		}

		// the child becomes the parent
		r = n
	}
}

// Delete returns true if the given word was removed from the trie.
// O(len(s))
func (t *Trie) Delete(s string) bool {
	n := traverse(t.start(s), s, t.s)

	// the word doesn't exist in the trie, so nothing to remove
	if n == nil || !n.e {
		return false
	}

	for n != nil {
		if n.nc == 0 {
			// the node has no children, so remove it
			if n.p != nil {
				n.p.nc--
				n.p.n[n.c] = nil
			}

			// move up, but only continue until
			// we find a terminating node
			n = n.p
			if n.e {
				break
			}
		} else {
			// the current node has children
			// in this case, the node is no longer terminating
			// but nothing can be deleted
			n.e = false
			break
		}
	}

	t.w--
	return true
}

// Has returns true if the trie contains the given word.
// O(len(s))
func (t *Trie) Has(s string) bool {
	n := traverse(t.start(s), s, t.s)

	return n != nil && n.e
}

// StartsWith returns all words in the trie that begin with
// the given string.
// O(N)
func (t *Trie) StartsWith(s string) (matches []string) {
	n := traverse(t.start(s), s, t.s)

	return append(matches, match(n, s)...)
}

// match recurisvely searches for matches for a given string.
func match(n *node, s string) (matches []string) {
	if n == nil {
		return matches
	}

	if n.e {
		matches = append(matches, s)
	}

	for _, c := range n.n {
		if c != nil {
			matches = append(matches, match(c, s+string(c.c))...)
		}
	}

	return
}

// traverse returns the last matching node for a given word.
func traverse(n *node, s string, size rune) *node {
	for i, c := range s {
		if n == nil || n.c != c {
			break
		}

		if i != len(s)-1 {
			n = n.next(rune(s[i+1]), size)
		}
	}

	return n
}

// Clear removes all the elements from the trie.
// O(1)
func (t *Trie) Clear() {
	t.Init(t.s)
	t.w = 0
}

// Len returns the number of words in the trie.
// O(1)
func (t *Trie) Len() int {
	return t.w
}

// start returns the first node under the root based on the
// word's first character and the trie's alphabet.
func (t *Trie) start(s string) *node {
	return t.r.n[(rune(s[0]) % t.s)]
}

// next returns the next node under the current node based
// on the given letter in the word and the trie's alphabet.
func (n *node) next(r rune, size rune) *node {
	return n.n[r%size]
}
