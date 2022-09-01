package ch24_kth_largest_element

type solution1UsingQuickSort struct {
}

// [5, 3, 1, 6, 4, 2] == sorted ==> [1, 2, 3, 4, 5, 6]
// 5th largest number: 2
//	1, 2, 5, 3, 6, 4] => [1] and [5, 3, 6, 4]
// 1st largest number: 6
// 2nd largest number: 5
//	1, 2, 5, 3, 6, 4] => [1] and [5, 3, 6, 4]
//	[5, 3, 6, 4], 2 => [3, 4, 6, 5], 1
//	[6, 5], 2 => [5, 6]
func (s1 solution1UsingQuickSort) findKthLargestElement(k int, arr []int) int {
	return quickSelect(arr, 0, len(arr)-1, len(arr)-k)
}

func quickSelect(arr []int, start, end int, largestElmtIdx int) int {
	if len(arr) <= 1 {
		return arr[0]
	}

	p := findPivotElementsPosition(arr, start, end)
	if largestElmtIdx == p {
		return arr[p]
	} else if largestElmtIdx < p { // left side of Pivot has the largest number
		return quickSelect(arr, start, p-1, largestElmtIdx)
	} else {
		return quickSelect(arr, p+1, end, largestElmtIdx)
	}
}

func findPivotElementsPosition(arr []int, start, end int) int {
	pivotIdx := end // last element position is pivot position

	i, j := start, start
	for j < pivotIdx {
		if arr[j] < arr[pivotIdx] {
			arr[i], arr[j] = arr[j], arr[i] // move j to the first lowest position
			i++
			j++
		} else {
			j++
		}
	}

	arr[i], arr[pivotIdx] = arr[pivotIdx], arr[i]

	return i
}
