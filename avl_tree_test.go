package algorithm

import (
	"fmt"
	"testing"
)

func TestAVLTree(t *testing.T) {
	avl := NewAVLTree()

	//inits := []int{40, 20, 30, 60, 10, 50, 70, 25, 15}
	inits := []int{18, 14, 20, 12, 16, 15}
	for _, v := range inits {
		ok := avl.Insert(v)

		t.Log(fmt.Sprintf("Insert %v: %v", v, avl.String()))
		if !ok {
			t.Errorf("Insert failed %v", v)
		}

		if !avl.Check() {
			t.Errorf("Insert invalid AVL tree.")
		}
	}

}
