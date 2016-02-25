package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()

	inits := []int{40, 20, 30, 60, 10, 50, 25, 15, 70}
	for _, v := range inits {
		ok := bst.Insert(v)

		t.Log(fmt.Sprintf("Insert %v: %v", v, bst.String()))
		if !ok {
			t.Errorf("Insert failed %v", v)
		}
	}

	finds := []int{15, 10, 35, 50}
	findResults := []bool{true, true, false, true}

	for i, v := range finds {
		_, ok := bst.Find(v)
		if ok != findResults[i] {
			t.Errorf("Find failed %v", v)
		}
	}

	dels := []int{20, 15, 65, 50, 25}
	delResults := []bool{true, true, false, true, true}

	for i, v := range dels {
		ok := bst.Delete(v)

		t.Log(fmt.Sprintf("Delete %v: %v", v, bst.String()))
		if ok != delResults[i] {
			t.Errorf("Delete failed %v", v)
		}
	}
}
