// Package trie implements a trie.
package trie

// T is the internal representation of a trie.
type T struct {
	root  *node
	words int
	size  rune
}

// node is the internal representation of a trie node.
type node struct {
	char   rune
	nodes  map[rune]*node
	end    bool
	parent *node
	value  interface{}
}

// Init initializes a trie with a given alphabet size.
// A trie must be initialized before it can be used.
// O(1)
func (t *T) Init(size rune) {
	if size < 1 {
		panic("Trie size must be a positive number")
	}

	t.root = &node{nodes: make(map[rune]*node, size)}
	t.size = size
}

// Insert adds a new word to the trie.
// The word may be accompanied by a value.
// Average: O(log(n)) Worst: O(n)
func (t *T) Insert(s string, v interface{}) {
	r := t.root

	for i, c := range s {
		n := r.next(c, t.size)
		end := i == len(s)-1
		if n == nil {
			n = &node{c, make(map[rune]*node, t.size), end, r, v}
			r.nodes[c%t.size] = n
			// increment the number of children for the parent
			// this information is useful for deletions
			if end {
				t.words++
			}
		} else if end && !n.end {
			// if the node previously existed, but wasn't a terminating string,
			// we need to now mark it as such (i.e. insert("foobar"), insert("foo"))
			// also increment the number of words in the trie in this scenario
			n.end = true
			n.value = v
			t.words++
		}

		// the child becomes the parent
		r = n
	}
}

// Delete returns true if the given word was removed from the trie.
// Average: O(log(n)) Worst: O(n)
func (t *T) Delete(s string) bool {
	n := traverse(t.start(s), s, t.size)

	// the word doesn't exist in the trie, so nothing to remove
	if n == nil || !n.end {
		return false
	}

	for n != nil {
		if len(n.nodes) == 0 {
			// the node has no children, so remove it
			if n.parent != nil {
				n.parent.nodes[n.char] = nil
			}

			// move up, but only continue until
			// we find a terminating node
			n = n.parent
			if n.end {
				break
			}
		} else {
			// the current node has children
			// in this case, the node is no longer terminating
			// but nothing can be deleted
			n.end = false
			break
		}
	}

	t.words--
	return true
}

// Has returns true if the trie contains the given word.
// Average: O(log(n)) Worst: O(n)
func (t *T) Has(s string) bool {
	_, r := t.Get(s)
	return r
}

// Get returns the value stored with the string and
// true if the trie contains the given word.
// Average: O(log(n)) Worst: O(n)
func (t *T) Get(s string) (interface{}, bool) {
	n := traverse(t.start(s), s, t.size)

	if n == nil || !n.end {
		return nil, false
	}

	return n.value, true
}

// StartsWith returns all words in the trie that begin with
// the given string.
// O(n)
func (t *T) StartsWith(s string) (matches []string) {
	n := traverse(t.start(s), s, t.size)

	return append(matches, match(n, s)...)
}

// match recurisvely searches for matches for a given string.
func match(n *node, s string) (matches []string) {
	if n == nil {
		return matches
	}

	if n.end {
		matches = append(matches, s)
	}

	for _, c := range n.nodes {
		if c != nil {
			matches = append(matches, match(c, s+string(c.char))...)
		}
	}

	return
}

// traverse returns the last matching node for a given word.
func traverse(n *node, s string, size rune) *node {
	for i, c := range s {
		if n == nil || n.char != c {
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
func (t *T) Clear() {
	t.Init(t.size)
	t.words = 0
}

// Len returns the number of words in the trie.
// O(1)
func (t *T) Len() int {
	return t.words
}

// start returns the first node under the root based on the
// word's first character and the trie's alphabet.
func (t *T) start(s string) *node {
	return t.root.nodes[(rune(s[0]) % t.size)]
}

// next returns the next node under the current node based
// on the given letter in the word and the trie's alphabet.
func (n *node) next(r rune, size rune) *node {
	return n.nodes[r%size]
}
