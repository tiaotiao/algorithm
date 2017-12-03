package algorithm

import "testing"

func TestQuickSort(t *testing.T) {
	cases := [][]int{
		{3},
		{},
		{2, 4, 6, 88, 2, 1, 0},
		{95, 4, 5},
		{-9, 4, 52, 0, -72, 6},
	}

	check := func(a []int) bool {
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				return false
			}
		}
		return true
	}

	for i, c := range cases {
		QuickSort(c)
		if !check(c) {
			t.Errorf("%d: %v", i, c)
		}
	}
}
