// Package trie implements a trie.
package trie

// Trie is the internal representation of a trie.
type Trie struct {
	r *node
	c int
	s rune
}

// node is the internal representation of a trie node.
type node struct {
	c rune
	n []*node
	e bool
}

// Init initializes a trie with a given alphabet size.
// A trie must be initialized before it can be used.
// O(1)
func (t *Trie) Init(size rune) {
	if size < 1 {
		panic("Trie size must be a positive number")
	}

	t.r = &node{n: make([]*node, size)}
	t.s = size
}

// Adds a new string to the trie.
// O(len(s))
func (t *Trie) Insert(s string) {
	r := []rune(s)
	t.r.n[t.first(r)] = insert(t.start(r), r, 0, t)
}

// insert recursively walks trie and adds nodes.
func insert(n *node, s []rune, i int, t *Trie) (r *node) {
	end := i == len(s)-1
	if r = n; n == nil {
		r = &node{s[i], make([]*node, t.s), end}
		// if this node is the string termination, incrase the count
		// so we can acurately return the number of words in the trie
		if end {
			t.c++
		}
	}

	if !end {
		i++
		r.n[s[i]%t.s] = insert(r.next(s[i], t.s), s, i, t)
	} else if !r.e {
		// if the node previously existed, but wasn't a terminating string,
		// we need to now mark it as such (i.e. insert("foobar"), insert("foo"))
		// also increment the number of words in the trie in this scenario
		r.e = true
		t.c++
	}

	return
}

// Clear removes all the elements from the trie.
// O(1)
func (t *Trie) Clear() {
	t.Init(t.s)
	t.c = 0
}

// Len returns the number of words in the trie.
// O(1)
func (t *Trie) Len() int {
	return t.c
}

// Has returns true if the trie contains the given word.
// O(len(s))
func (t *Trie) Has(s string) bool {
	r := []rune(s)
	n := traverse(t.start(r), r, t.s)

	return n != nil && n.e
}

// StartsWith returns all words in the trie that begin with
// the given string.
// O(log(N))
func (t *Trie) StartsWith(s string) (matches []string) {
	r := []rune(s)
	n := traverse(t.start(r), r, t.s)

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
func traverse(n *node, r []rune, s rune) *node {
	for i, c := range r {
		if n == nil || n.c != c {
			break
		}

		if i != len(r)-1 {
			n = n.next(r[i+1], s)
		}
	}

	return n
}

// first returns the first character in the word based
// on the size of the trie's alphabet.
func (t *Trie) first(r []rune) rune {
	return r[0] % t.s
}

// start returns the first node under the root based on the
// word's first character and the trie's alphabet.
func (t *Trie) start(r []rune) *node {
	return t.r.n[t.first(r)]
}

// next returns the next node under the current node based
// on the given letter in the word and the trie's alphabet.
func (n *node) next(r rune, s rune) *node {
	return n.n[r%s]
}
