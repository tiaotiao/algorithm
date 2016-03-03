package algorithm

import (
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
		//println(bt.String())
	}
}
