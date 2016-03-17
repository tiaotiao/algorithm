package algorithm

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tiaotiao/go/util"
)

type BTree struct {
	root   *bnode
	degree int // the max number of children
	most   int // the max number of elements
	least  int // the min number of elements

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
	t.most = degree - 1
	t.least = t.most / 2
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

	pos, exist := t.locate(n, v)
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
	pos, exist := t.locate(n, v)

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
	t.add(n, pos, elem, child, "right")

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

func (t *BTree) Delete(v interface{}) (interface{}, bool) {
	val, ok := t.delete(t.root, v)
	if t.root != nil && len(t.root.elements) == 0 {
		// reduce the height of the tree
		t.root = t.root.children[0]
	}
	return val, ok
}

func (t *BTree) delete(n *bnode, v interface{}) (interface{}, bool) {
	if n == nil {
		return nil, false
	}

	pos, exist := t.locate(n, v)
	if !exist {
		deleted, ok := t.delete(n.children[pos], v)
		t.adjust(n, pos)
		return deleted, ok
	}

	// this node may in the middle of the tree.
	// so delete the max element from left subtree instead, and then replace the element.
	deleted := n.elements[pos]
	maxLeft := t.deleteMax(n.children[pos]) // delete the max value from left the child
	if maxLeft != nil {
		n.elements[pos] = maxLeft
		t.adjust(n, pos)
		return deleted, true
	}

	// this is a leaf node.
	// so just delete the element here.
	deleted = n.elements[pos]
	t.del(n, pos, "right") // all children are nil, so it doesn't matter left or right

	return deleted, true
}

func (t *BTree) deleteMax(n *bnode) interface{} {
	if n == nil {
		return nil
	}

	pos := len(n.children) - 1
	child := n.children[pos]
	if child != nil {
		deleted := t.deleteMax(child)
		t.adjust(n, pos)
		return deleted
	}

	deleted, _ := t.pop(n, "right")
	return deleted
}

func (t *BTree) adjust(parent *bnode, ch int) {
	if parent == nil {
		return
	}
	child := parent.children[ch]
	if child == nil {
		return
	}

	if !t.tooFew(child) {
		return // no need to adjust
	}

	sibling, mid, side := t.chooseSibling(parent, ch)
	if sibling == nil {
		panic("sibling can not be nil")
	}

	if len(sibling.elements) > t.least {
		// borrow element from sibling
		t.borrow(parent, child, mid, sibling, side)
		return
	}

	// need to combine nodes
	t.combine(parent, child, mid, sibling, side)
}

func (t *BTree) combine(parent, child *bnode, mid int, sibling *bnode, sideOfSibling string) {
	// fmt.Printf("combine %v, %v, %v\n", child, mid, sibling)
	// remove the element in the middle of child and sibling from parent
	middle, _ := t.del(parent, mid, "right")
	// assume child is on the left
	if sideOfSibling == "left" {
		child, sibling = sibling, child
	}
	// combine nodes
	child.elements = append(child.elements, middle)
	child.elements = append(child.elements, sibling.elements...)
	child.children = append(child.children, sibling.children...)
}

func (t *BTree) borrow(parent, child *bnode, mid int, sibling *bnode, sideOfSibling string) {
	// fmt.Printf("borrow %v, %v, %v\n", child, mid, sibling)
	otherSide := t.otherSide(sideOfSibling)

	// borrow an element from sibling
	borrowedElem, borrowedChild := t.pop(sibling, otherSide)
	// take the element from parent in the middle of child and sibling
	middle := parent.elements[mid]
	// push to child
	t.push(child, middle, borrowedChild, sideOfSibling)
	// put borrowed element in the middle
	parent.elements[mid] = borrowedElem
}

func (t *BTree) chooseSibling(parent *bnode, ch int) (sibling *bnode, middle int, side string) {
	var left, right *bnode = nil, nil
	if ch-1 >= 0 {
		left = parent.children[ch-1]
	}
	if ch+1 < len(parent.children) {
		right = parent.children[ch+1]
	}
	if left == nil {
		return right, ch, "right"
	}
	if right == nil {
		return left, ch - 1, "left"
	}
	if len(left.elements) > len(right.elements) {
		return left, ch - 1, "left"
	}
	return right, ch, "right"
}

func (t *BTree) String() string {
	s, _ := t.toString(t.root, 0)
	return s
}

func (t *BTree) toString(n *bnode, depth int) (str string, r int) {
	if n == nil {
		return "", 0
	}
	var s string
	for i, c := range n.children {
		s, r = t.toString(c, depth+1)
		if s != "" {
			s = fmt.Sprintf("[%s]", s)
		}
		// if r == 1 {
		// 	s = fmt.Sprintf("%d:", depth+1) + s
		// }
		str += s
		if i < len(n.elements) {
			se := fmt.Sprintf("(%v)", n.elements[i])
			// if r > 0 {
			// 	se = fmt.Sprintf("%d:", depth) + se
			// }
			sp := strings.Repeat(" ", r)
			se = sp + se + sp
			str += se
		}
	}
	return fmt.Sprintf("%s", str), r + 1
}

func (t *BTree) Check() error {
	_, _, _, ok := t.check(t.root, 0)
	if !ok {
		return fmt.Errorf("invalid b-tree")
	}
	return nil
}

func (t *BTree) check(n *bnode, depth int) (height int, min, max interface{}, ok bool) {
	if n == nil {
		return 0, nil, nil, true
	}

	if t.tooMany(n) {
		return -1, nil, nil, false
	}

	if depth > 0 {
		if t.tooFew(n) {
			return -1, nil, nil, false
		}
	}

	height = -1
	min, max = nil, nil
	for idx, ch := range n.children {
		h, childmin, childmax, ok := t.check(ch, depth+1)
		if !ok {
			return -1, nil, nil, false
		}

		if childmin != nil {
			if min == nil {
				min = childmin
			}
			if max != nil {
				if t.Compare(max, childmin) > 0 {
					return -1, nil, nil, false
				}
			}
		}
		if childmax != nil {
			max = childmax
		}

		if idx < len(n.elements) {
			elem := n.elements[idx]
			if min == nil {
				min = elem
			}

			if max != nil {
				if t.Compare(max, elem) > 0 {
					return -1, nil, nil, false
				}
			}
			max = elem
		}

		if height == -1 {
			height = h
		} else if height != h {
			return -1, nil, nil, false
		}
	}

	return height + 1, min, max, true
}

func (t *BTree) newNode() *bnode {
	n := new(bnode)
	n.elements = make([]interface{}, 0, t.degree)
	n.children = make([]*bnode, 0, t.degree+1)
	return n
}

func (t *BTree) locate(n *bnode, v interface{}) (int, bool) {
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

func (t *BTree) otherSide(s string) string {
	if s == "left" {
		return "right"
	}
	return "left"
}

// Add the elem to this node
func (t *BTree) add(n *bnode, pos int, elem interface{}, child *bnode, side string) {
	ch := pos // left child position
	if side == "right" {
		ch = pos + 1
	}
	util.Insert(&n.elements, pos, elem)
	util.Insert(&n.children, ch, child)
}

func (t *BTree) push(n *bnode, elem interface{}, child *bnode, side string) {
	// fmt.Printf("push %v, %v, %v\n", n, elem, side)
	if side == "left" {
		t.add(n, 0, elem, child, "left")
	} else {
		t.add(n, len(n.elements), elem, child, "right")
	}
}

func (t *BTree) pop(n *bnode, side string) (interface{}, *bnode) {
	// fmt.Printf("pop %v, %v\n", n, side)
	if side == "left" {
		return t.del(n, 0, "left")
	} else {
		return t.del(n, len(n.elements)-1, "right")
	}
}

func (t *BTree) del(n *bnode, pos int, side string) (interface{}, *bnode) {
	ch := pos // left child position
	if side == "right" {
		ch = pos + 1
	}
	deleted := n.elements[pos]
	child := n.children[ch]
	util.Remove(&n.elements, pos)
	util.Remove(&n.children, ch)
	return deleted, child
}

func (t *BTree) tooFew(n *bnode) bool {
	return len(n.elements) < t.least
}

func (t *BTree) tooMany(n *bnode) bool {
	return len(n.elements) > t.most
}
