package algorithm

import (
	"fmt"
	"testing"
)

func TestAVLTree(t *testing.T) {
	avl := NewAVLTree()

	inits := []int{40, 20, 30, 60, 10, 50, 70, 25, 15}
	//inits := []int{18, 14, 20, 12, 16, 15}
	for _, v := range inits {
		ok := avl.Insert(v)

		t.Log(fmt.Sprintf("Insert %v: %v", v, avl.String()))
		if !ok {
			t.Errorf("Insert failed %v", v)
		}

		if !avl.Check() {
			t.Errorf("Check failed %v", avl.String())
		}
	}

	finds := []int{15, 10, 35, 50}
	findResults := []bool{true, true, false, true}

	for i, v := range finds {
		_, ok := avl.Find(v)
		if ok != findResults[i] {
			t.Errorf("Find failed %v", v)
		}
	}

	dels := []int{20, 15, 65, 50, 25}
	delResults := []bool{true, true, false, true, true}

	for i, v := range dels {
		ok := avl.Delete(v)

		t.Log(fmt.Sprintf("Delete %v: %v", v, avl.String()))
		if ok != delResults[i] {
			t.Errorf("Delete failed %v", v)
		}

		if !avl.Check() {
			t.Errorf("Check failed %v", avl.String())
		}
	}
}
