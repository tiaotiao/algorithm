package algorithm

func BinarySearch(a []int, val int) int {
	if len(a) <= 0 {
		return -1
	}
	left := 0
	right := len(a) - 1
	mid := (left + right) / 2

	// fmt.Printf("Bearch %v, len=%v, mid=%v, v=%v\n", a, len(a), mid, val)

	if a[mid] == val {
		if left == mid {
			return mid
		}
		return BinarySearch(a[:mid+1], val) // find the left most one if duplicated
	} else if val < a[mid] {
		return BinarySearch(a[:mid], val)
	} else {
		ret := BinarySearch(a[mid+1:], val)
		if ret == -1 {
			return -1
		}
		return mid + 1 + ret
	}
}
