package algorithm

import "testing"

func TestKMP(t *testing.T) {
	cases := []struct {
		s, p string
		e    int
	}{
		{s: "aaaabab", p: "cabd", e: -1},
		{s: "aaab", p: "aaabaab", e: -1},
		{s: "aaaabab", p: "aaba", e: 2},
		{s: "aababaabaabaabaabab", p: "aabaabab", e: 11},
	}

	for i, c := range cases {
		r := KMP(c.s, c.p)
		if r != c.e {
			t.Errorf("KMP failed %d: s=%v, p=%v, expect=%v, got=%v", i, c.s, c.p, c.e, r)
		}
	}
}
