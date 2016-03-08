package algorithm

import (
	"fmt"
	"github.com/tiaotiao/go/util"
	"strings"
	"testing"
)

func TestBTree(t *testing.T) {
	bt := NewBTree(5)
	bt.Compare = func(a, b interface{}) int {
		x := a.(string)
		y := b.(string)
		return strings.Compare(x, y)
	}

	insertions := []string{"C", "N", "G", "A", "H", "E", "K", "Q", "M", "F", "W", "L", "T", "Z", "D", "P", "R", "X", "Y", "S"}

	for _, v := range insertions {
		bt.Insert(v)
		t.Log(bt.String())

		if !bt.Check() {
			t.Errorf("Insert failed %v, %v", v, bt.String())
		}
	}

	finds := []string{"C", "D", "P", "J", "N", "B"}

	for _, v := range finds {
		_, ok := bt.Find(v)

		exists := (util.Find(insertions, v) >= 0)

		if ok != exists {
			t.Error(fmt.Sprintf("Find error: %v, should be %v", v, exists))
		}
	}
}
