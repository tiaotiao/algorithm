package algorithm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tiaotiao/go/util"
)

func TestBTree(t *testing.T) {
	bt := NewBTree(5)
	bt.Compare = func(a, b interface{}) int {
		x := a.(string)
		y := b.(string)
		return strings.Compare(x, y)
	}

	// insert
	insertions := []string{"C", "N", "G", "A", "H", "E", "K", "Q", "M", "F", "W", "L", "T", "Z", "D", "P", "R", "X", "Y", "S"}

	for _, v := range insertions {
		bt.Insert(v)
		t.Log(fmt.Sprintf("Insert %v: %v", v, bt.String()))

		if bt.Check() != nil {
			t.Errorf("Insert failed %v, %v", v, bt.String())
		}
	}

	// find
	finds := []string{"C", "D", "P", "J", "N", "B"}

	for _, v := range finds {
		_, ok := bt.Find(v)

		exists := (util.Find(insertions, v) >= 0)

		if ok != exists {
			t.Error(fmt.Sprintf("Find error: %v, should be %v", v, exists))
		}
	}

	// delete
	deletion := []string{"H", "T", "R", "E"}

	for _, v := range deletion {
		_, ok := bt.Delete(v)

		t.Log(fmt.Sprintf("Delete %v: %v", v, bt.String()))

		if !ok {
			t.Error(fmt.Sprintf("Delete error: %v, %v", v, bt.String()))
		}

		if bt.Check() != nil {
			t.Errorf("Delete failed %v, %v", v, bt.String())
		}
	}
}
