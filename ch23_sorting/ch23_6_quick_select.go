package ch23_sorting

type QuickSelect struct {
}

func (s QuickSelect) smallestElement(arr []int, k int) int {
	myIdx := len(arr) - k
	selectRecursive(arr, 0, myIdx)
	return arr[myIdx]
}

func selectRecursive(arr []int, startIdx, myIdx int) {
	p := sortPivot(arr)

	selectIdx := startIdx + p

	if myIdx < selectIdx {
		left := arr[:p]
		if len(left) > 1 {
			selectRecursive(left, selectIdx, myIdx)
		}
	} else {
		right := arr[p+1:]
		if len(right) > 1 {
			selectRecursive(right, selectIdx, myIdx)
		}
	}
}
