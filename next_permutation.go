package algorithm

import (
	"github.com/tiaotiao/go/util"
)

func NextPermutation(perm []int) bool {
	n := len(perm)

	k := -1
	for i := n - 1; i > 0; i-- {
		if perm[i-1] < perm[i] {
			k = i - 1
			break
		}
	}

	if k == -1 {
		return false
	}

	for i := n - 1; i > k; i-- {
		if perm[i] > perm[k] {
			perm[i], perm[k] = perm[k], perm[i]
			break
		}
	}

	util.Reverse(perm[k+1 : n])

	return true
}
