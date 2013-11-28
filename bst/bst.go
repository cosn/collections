// Package bst implements an unbalanced binary search tree
package bst

import "fmt"

// BST is the internal representation of a binary search tree
type BST struct {
	r *node
	c int
}

// node is the internal representation of a binary tree node
type node struct {
	k    int
	v    interface{}
	l, r *node
}

// TraversalType represents one of the three know traversals
type TraversalType int

const (
	InOrder TraversalType = iota
	PreOrder
	PostOrder
)

// Insert adds a given key+value to the tree and returns true if it was added
// Average: O(log(n)) Worst: O(n)
func (t *BST) Insert(k int, v interface{}) (added bool) {
	t.r, added = insert(t.r, k, v)
	if added {
		t.c++
	}

	return
}

// insert recusively adds a key+value in the tree
func insert(n *node, k int, v interface{}) (r *node, added bool) {
	if r = n; n == nil {
		// keep track of how many elements we have in the tree
		// to optimize the channel length during traversal
		r = &node{k: k, v: v}
		added = true
	} else if k < n.k {
		r.l, added = insert(n.l, k, v)
	} else if k > n.k {
		r.r, added = insert(n.r, k, v)
	}

	return
}

// Delete removes a given key from the tree and returns true if it was removed
// Average: O(log(n)) Worst: O(n)
func (t *BST) Delete(k int) (deleted bool) {
	_, deleted = delete(t.r, k)
	if deleted {
		t.c--
	}

	return
}

// delete recursively deletes a key from the tree
func delete(n *node, k int) (r *node, deleted bool) {
	if r = n; n == nil {
		return nil, false
	}

	if k < n.k {
		r.l, deleted = delete(n.l, k)
	} else if k > n.k {
		r.r, deleted = delete(n.r, k)
	} else {
		if n.l != nil && n.r != nil {
			// find the right most element in the left subtree
			s := n.l
			for s.r != nil {
				s = s.r
			}
			r.k = s.k
			r.v = s.v
			r.l, deleted = delete(s, s.k)
		} else if n.l != nil {
			r = n.l
			deleted = true
		} else if n.r != nil {
			r = n.r
			deleted = true
		} else {
			r = nil
			deleted = true
		}
	}

	return
}

// Find returns the value found at the given key
// Average: O(log(n)) Worst: O(n)
func (t *BST) Find(k int) interface{} {
	return find(t.r, k)
}

func find(n *node, k int) interface{} {
	if n == nil {
		return nil
	}

	if n.k == k {
		return n.v
	} else if k < n.k {
		return find(n.l, k)
	} else if k > n.k {
		return find(n.r, k)
	}

	return nil
}

// Clear removes all the nodes from the tree
// O(n)
func (t *BST) Clear() {
	t.r = clear(t.r)
	t.c = 0
}

// clear recursively removes all the nodes
func clear(n *node) *node {
	if n != nil {
		n.l = clear(n.l)
		n.r = clear(n.r)
	}
	n = nil

	return n
}

// Traverse provides an iterator over the tree
// O(n)
func (t *BST) Traverse(tt TraversalType) <-chan interface{} {
	c := make(chan interface{}, t.c)
	go func() {
		switch tt {

		case InOrder:
			inOrder(t.r, c)
		case PreOrder:
			preOrder(t.r, c)
		case PostOrder:
			postOrder(t.r, c)
		}
		close(c)
	}()

	return c
}

// inOrder returns the left, parent, right nodes
func inOrder(n *node, c chan interface{}) {
	if n == nil {
		return
	}

	inOrder(n.l, c)
	c <- n.v
	inOrder(n.r, c)
}

// preOrder returns the parent, left, right nodes
func preOrder(n *node, c chan interface{}) {
	if n == nil {
		return
	}

	c <- n.v
	preOrder(n.l, c)
	preOrder(n.r, c)
}

// postOrder returns the left, right, parent nodes
func postOrder(n *node, c chan interface{}) {
	if n == nil {
		return
	}

	postOrder(n.l, c)
	postOrder(n.r, c)
	c <- n.v
}

// String prints the nodes in the tree
func (t *BST) String() (s string) {
	print(t.r, &s)
	return
}

// print recusively prints the pre-order nodes
func print(n *node, s *string) {
	if n == nil {
		return
	}

	*s += fmt.Sprintf("%p %v\n", n, n)
	print(n.l, s)
	print(n.r, s)
}
