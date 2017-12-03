package algorithm

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		a []int
		v int
		e int
	}{
		{a: []int{1, 2, 2, 5, 6}, v: 5, e: 3},
		{a: []int{1, 2, 2, 5, 6}, v: 3, e: -1},
		{a: []int{1, 2, 2, 5, 6}, v: 2, e: 1}, // find the first one
		{a: []int{1, 1, 1, 5, 6}, v: 1, e: 0}, // find the first one
	}

	for i, c := range cases {
		r := BinarySearch(c.a, c.v)
		if r != c.e {
			t.Errorf("%d: a=%v, v=%v, expect=%v, got=%v", i, c.a, c.v, c.e, r)
		}
	}
}
