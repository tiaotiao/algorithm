package algorithm

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a int, b ...int) int {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}

func min(a int, b ...int) int {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
