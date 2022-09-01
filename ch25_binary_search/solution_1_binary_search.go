package ch25_binary_search

type solution1BinarySearch struct {
}

func (s1 solution1BinarySearch) search(arr []int, val int) int {
	pL, pR := 0, len(arr)
	for pL <= pR {
		mid := (pL + pR) / 2
		if arr[mid] == val {
			return mid
		} else if val < arr[mid] {
			pR = mid - 1
		} else {
			pL = mid + 1
		}
	}

	return -1
}
