package algorithm

type AVLTree struct {
	*BinarySearchTree
}

func NewAVLTree() *AVLTree {
	t := new(AVLTree)
	t.BinarySearchTree = NewBinarySearchTree()
	t.insert = t._insert
	t.delete = t._delete
	t.del_min = t._del_min
	return t
}

func (t *AVLTree) _insert(p **node, v interface{}) bool {
	ok := t.BinarySearchTree._insert(p, v)
	if !ok {
		return false
	}

	t.rebalance(p)

	return true
}

func (t *AVLTree) _delete(p **node, v interface{}) bool {
	ok := t.BinarySearchTree._delete(p, v)
	if !ok {
		return false
	}

	t.rebalance(p)

	return true
}

func (t *AVLTree) _del_min(p **node) *node {
	n := t.BinarySearchTree._del_min(p)

	t.rebalance(p)

	return n
}

func (t *AVLTree) rebalance(p **node) bool {
	if *p == nil {
		return false
	}

	defer t.updateHeight(*p)

	leftHeight := t.height((*p).left)
	rightHeight := t.height((*p).right)

	if abs(leftHeight-rightHeight) <= 1 {
		return false
	}

	// need to rebanlance
	if leftHeight > rightHeight {
		leftChild := (*p).left
		if t.height(leftChild.left) < t.height(leftChild.right) {
			t.rotateLeft(&(*p).left)
		}
		t.rotateRight(p)

	} else {
		rightChild := (*p).right
		if t.height(rightChild.left) > t.height(rightChild.right) {
			t.rotateRight(&(*p).right)
		}
		t.rotateLeft(p)
	}

	return true
}

func (t *AVLTree) rotateRight(p **node) bool {
	if *p == nil {
		return false
	}
	leftChild := (*p).left
	if leftChild == nil {
		return false
	}

	(*p).left = leftChild.right
	leftChild.right = *p
	*p = leftChild

	t.updateHeight((*p).right)
	t.updateHeight(*p)

	return true
}

func (t *AVLTree) rotateLeft(p **node) bool {
	if *p == nil {
		return false
	}
	rightChild := (*p).right
	if rightChild == nil {
		return false
	}

	(*p).right = rightChild.left
	rightChild.left = *p
	*p = rightChild

	t.updateHeight((*p).left)
	t.updateHeight(*p)

	return true
}

// To check if this is a valid AVL tree
func (t *AVLTree) Check() bool {
	_, ok := t.check(t.root)
	return ok
}

func (t *AVLTree) check(n *node) (height int, valid bool) {
	if n == nil {
		return 0, true
	}
	leftHeight, leftOK := t.check(n.left)
	rightHeight, rightOK := t.check(n.right)
	if !leftOK || !rightOK {
		return 0, false
	}
	if abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}
	height = max(leftHeight, rightHeight) + 1
	return height, true
}

func (t *AVLTree) height(n *node) int {
	if n == nil {
		return 0
	}
	if n.extra == nil {
		return 0
	}
	return n.extra.(int)
}

func (t *AVLTree) updateHeight(n *node) int {
	if n == nil {
		return 0
	}
	leftHeight := t.height(n.left)
	rightHeight := t.height(n.right)

	height := max(leftHeight, rightHeight) + 1
	n.extra = height
	return height
}
