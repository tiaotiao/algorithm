package algorithm

import (
	"fmt"
	"github.com/tiaotiao/go/util"
	"sort"
	"strings"
)

type BTree struct {
	root   *bnode
	degree int

	Compare func(a, b interface{}) int
}

// A node of degree 4 is like this:
// elements:	[0] [1] [2]
//				/  |   |  \
// children:  [0] [1] [2] [3]
type bnode struct {
	elements []interface{} // [0 ~ degree-1]
	children []*bnode      // [0 ~ degree]
}

func NewBTree(degree int) *BTree {
	t := new(BTree)
	t.root = nil
	t.degree = degree
	t.Compare = CompareInt
	return t
}

func New234Tree() *BTree {
	return NewBTree(4)
}

func (t *BTree) Find(v interface{}) (interface{}, bool) {
	return t.find(t.root, v)
}

func (t *BTree) find(n *bnode, v interface{}) (interface{}, bool) {
	if n == nil {
		return nil, false
	}

	pos, exist := t.searchNode(n, v)
	if exist {
		return n.elements[pos], true
	}

	return t.find(n.children[pos], v)
}

func (t *BTree) Insert(v interface{}) {
	elem, child := t._insert(t.root, v)
	if elem != nil {
		newnode := t.newNode()
		newnode.elements = append(newnode.elements, elem)
		newnode.children = append(newnode.children, t.root, child)
		t.root = newnode
	}
}

func (t *BTree) _insert(n *bnode, v interface{}) (elem interface{}, child *bnode) {
	if n == nil {
		return v, nil
	}

	// find proper a place for v
	pos, exist := t.searchNode(n, v)

	if exist { // the value is found in this node, replace it
		n.elements[pos] = v
		return nil, nil
	}

	// insert into subtree
	elem, child = t._insert(n.children[pos], v)
	if elem == nil {
		// already done
		return nil, nil
	}

	// Add the elem to this node.
	// This elem may be the value intended to insert if subtree is nil.
	// Or may be come from the split of subtree.
	util.Insert(&n.elements, pos, elem)
	util.Insert(&n.children, pos+1, child)

	if len(n.elements) < t.degree { // Not full, done.
		return nil, nil
	}

	// This node is full, split into two nodes
	idx := len(n.elements) / 2 // ceil(len/2)
	mid := n.elements[idx]

	newnode := t.newNode()
	newnode.elements = append(newnode.elements, n.elements[idx+1:]...)
	newnode.children = append(newnode.children, n.children[idx+1:]...)

	n.elements = n.elements[:idx]
	n.children = n.children[:idx+1]

	// Return splited element to parent.
	return mid, newnode
}

func (t *BTree) String() string {
	s, _ := t.toString(t.root)
	return s
}

func (t *BTree) toString(n *bnode) (str string, r int) {
	if n == nil {
		return "", 0
	}
	var s string
	for i, c := range n.children {
		s, r = t.toString(c)
		str += fmt.Sprintf("[%s]", s)
		if i < len(n.elements) {
			se := fmt.Sprintf("(%v)", n.elements[i])
			sp := strings.Repeat(" ", r)
			se = sp + se + sp
			str += se
		}
	}
	return fmt.Sprintf("%s", str), r + 1
}

func (t *BTree) newNode() *bnode {
	n := new(bnode)
	n.elements = make([]interface{}, 0, t.degree)
	n.children = make([]*bnode, 0, t.degree+1)
	return n
}

func (t *BTree) searchNode(n *bnode, v interface{}) (int, bool) {
	exist := false
	pos := sort.Search(len(n.elements), func(i int) bool {
		c := t.Compare(n.elements[i], v)
		if c == 0 {
			exist = true
		}
		return c >= 0
	})
	return pos, exist
}
