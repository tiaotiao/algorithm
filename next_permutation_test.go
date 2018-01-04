package algorithm

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {

	cases := []struct {
		p  []int
		q  []int
		ok bool
	}{
		{
			p:  []int{1, 2, 4, 5, 5, 3},
			q:  []int{1, 2, 5, 3, 4, 5},
			ok: true,
		},
		{
			p:  []int{7, 5, 3, 2},
			ok: false,
		},
		{
			p:  []int{4, 8, 5, 6, 3, 1},
			q:  []int{4, 8, 6, 1, 3, 5},
			ok: true,
		},
		{
			p:  []int{3, 6, 6},
			q:  []int{6, 3, 6},
			ok: true,
		},
		{
			p:  []int{1, 2},
			q:  []int{2, 1},
			ok: true,
		},
		{
			p:  []int{2},
			ok: false,
		},
	}

	for i, c := range cases {
		b := NextPermutation(c.p)
		if b != c.ok {
			t.Errorf("%v: %v, b %v != %v", i, c.p, b, c.ok)
			continue
		}
		if b == false {
			continue
		}
		if !reflect.DeepEqual(c.p, c.q) {
			t.Errorf("%v: %v, expect=%v", i, c.p, c.q)
		}
	}
}
