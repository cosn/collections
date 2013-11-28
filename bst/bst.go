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
	v    int
	l, r *node
}

// TraversalType represents one of the three know traversals
type TraversalType int

const (
	InOrder TraversalType = iota
	PreOrder
	PostOrder
)

// Insert adds a given value to the tree and returns true if it was added
// Average: O(log(n)) Worst: O(n)
func (t *BST) Insert(v int) (added bool) {
	t.r, added = insert(t.r, v)
	if added {
		t.c++
	}

	return
}

// insert recusively adds a value in the tree
func insert(n *node, v int) (r *node, added bool) {
	if r = n; n == nil {
		// keep track of how many elements we have in the tree
		// to optimize the channel length during traversal
		r = &node{v: v}
		added = true
	} else if v < n.v {
		r.l, added = insert(n.l, v)
	} else if v > n.v {
		r.r, added = insert(n.r, v)
	}

	return
}

// Delete removes a given value from the tree and retruns true if it was removed
// Average: O(log(n)) Worst: O(n)
func (t *BST) Delete(v int) (deleted bool) {
	_, deleted = delete(t.r, v)
	if deleted {
		t.c--
	}

	return
}

// delete recursively deletes a value from the tree
func delete(n *node, v int) (r *node, deleted bool) {
	if r = n; n == nil {
		return nil, false
	}

	if v < n.v {
		r.l, deleted = delete(n.l, v)
	} else if v > n.v {
		r.r, deleted = delete(n.r, v)
	} else {
		if n.l != nil && n.r != nil {
			// find the right most element in the left subtree
			s := n.l
			for s.r != nil {
				s = s.r
			}
			r.v = s.v
			r.l, deleted = delete(s, s.v)
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

// Contains returns true if the value given exists in the tree
// Average: O(log(n)) Worst: O(n)
func (t *BST) Contains(v int) bool {
	return contains(t.r, v)
}

func contains(n *node, v int) bool {
	if n == nil {
		return false
	}

	if n.v == v {
		return true
	} else if v < n.v {
		return contains(n.l, v)
	} else if v > n.v {
		return contains(n.r, v)
	}

	return false
}

// Traverse provides an iterator over the tree
// O(n)
func (t *BST) Traverse(tt TraversalType) <-chan int {
	c := make(chan int, t.c)
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
func inOrder(n *node, c chan int) {
	if n == nil {
		return
	}

	inOrder(n.l, c)
	c <- n.v
	inOrder(n.r, c)
}

// preOrder returns the parent, left, right nodes
func preOrder(n *node, c chan int) {
	if n == nil {
		return
	}

	c <- n.v
	preOrder(n.l, c)
	preOrder(n.r, c)
}

// postOrder returns the left, right, parent nodes
func postOrder(n *node, c chan int) {
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
