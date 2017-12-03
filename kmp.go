package algorithm

func KMP(str, pattern string) int {
	n := len(str)
	m := len(pattern)
	if n < m {
		return -1
	}

	next := kmp_calc_next(pattern)

	i, j := 0, 0
	for i < n && j < m {
		if str[i] == pattern[j] {
			i += 1
			j += 1
			continue
		}
		if j == 0 {
			i += 1
			continue
		}
		j = next[j-1]
	}

	if j >= m {
		return i - m
	}
	return -1
}

func kmp_calc_next(pattern string) []int {
	if len(pattern) <= 0 {
		return nil
	}

	next := make([]int, len(pattern))
	next[0] = 0
	m := len(pattern)
	i, j := 1, 0

	for i < m && j < m {
		if pattern[i] == pattern[j] {
			next[i] = j + 1
			i += 1
			j += 1
			continue
		}
		if j == 0 {
			next[i] = 0
			i += 1
			continue
		}
		j = next[j-1]
	}

	return next
}
