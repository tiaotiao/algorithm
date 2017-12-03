package algorithm

func QuickSort(a []int) {
	if len(a) <= 1 {
		return
	}

	left := 0
	right := len(a) - 1
	x := a[left]

	for left < right {
		for left < right && x < a[right] {
			right--
		}
		if left < right {
			a[left] = a[right]
			left++
		}
		for left < right && x > a[left] {
			left++
		}
		if left < right {
			a[right] = a[left]
			right--
		}
	}
	a[left] = x

	QuickSort(a[:left])
	QuickSort(a[right+1:])
}
