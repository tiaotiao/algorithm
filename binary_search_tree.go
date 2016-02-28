package algorithm

import (
	"fmt"
)

type BinarySearchTree struct {
	root    *node
	compare func(a interface{}, b interface{}) int

	insert  func(p **node, v interface{}) bool
	delete  func(p **node, v interface{}) bool
	del_min func(p **node) *node
}

type node struct {
	value interface{}

	left  *node
	right *node

	extra interface{} // for extend
}

func NewBinarySearchTree() *BinarySearchTree {
	t := new(BinarySearchTree)
	t.root = nil
	t.compare = CompareInt
	t.insert = t._insert
	t.delete = t._delete
	t.del_min = t._del_min
	return t
}

func (t *BinarySearchTree) Insert(v interface{}) bool {
	return t.insert(&t.root, v)
}

func (t *BinarySearchTree) _insert(p **node, v interface{}) bool {
	if *p == nil {
		*p = t.newNode(v)
		return true
	}

	c := t.compare(v, (*p).value)

	if c == 0 {
		(*p).value = v
		return false
	} else if c < 0 {
		return t.insert(&((*p).left), v)
	} else {
		return t.insert(&((*p).right), v)
	}
}

func (t *BinarySearchTree) Find(v interface{}) (interface{}, bool) {
	n, ok := t.find(t.root, v)
	if !ok {
		return nil, false
	}
	return n.value, true
}

func (t *BinarySearchTree) find(n *node, v interface{}) (*node, bool) {
	if n == nil {
		return nil, false
	}

	c := t.compare(v, n.value)

	if c == 0 {
		return n, true
	} else if c < 0 {
		return t.find(n.left, v)
	} else {
		return t.find(n.right, v)
	}
}

func (t *BinarySearchTree) Delete(v interface{}) bool {
	return t.delete(&t.root, v)
}

func (t *BinarySearchTree) _delete(p **node, v interface{}) bool {
	n := *p
	if n == nil {
		return false
	}

	c := t.compare(v, n.value)

	if c == 0 { // delete current node
		t.del_node(p)
		return true
	} else if c < 0 {
		return t.delete(&n.left, v)
	} else {
		return t.delete(&n.right, v)
	}
}

func (t *BinarySearchTree) del_node(p **node) *node {
	if *p == nil {
		return nil
	}
	n := *p
	defer func() {
		n.right, n.left = nil, nil
	}()

	if n.right == nil { // if has no right subtree.
		*p = n.left // use the left subtree as replacement.
		return n
	}

	// use the minimun node from the right subtree as replacement.
	min := t.del_min(&(*p).right)
	min.left = n.left
	min.right = n.right
	(*p) = min
	return n
}

func (t *BinarySearchTree) _del_min(p **node) *node {
	if (*p).left != nil {
		return t.del_min(&(*p).left)
	}
	return t.del_node(p)
}

func (t *BinarySearchTree) String() string {
	return t.toString(t.root)
}

func (t *BinarySearchTree) toString(n *node) string {
	if n == nil {
		return ""
	}
	l := t.toString(n.left)
	r := t.toString(n.right)
	return fmt.Sprintf("[%s] %v [%s]", l, n.value, r)
}

func (t *BinarySearchTree) newNode(v interface{}) *node {
	n := new(node)
	n.value = v
	n.left = nil
	n.right = nil
	return n
}

func CompareInt(a interface{}, b interface{}) int {
	x := a.(int)
	y := b.(int)
	return x - y
}
