package ch26_find_start_and_end_index

type solution2BinarySearch struct {
}

func (s2 solution2BinarySearch) findStartingAndEndingIndex(arr []int, target int) []int {
	if len(arr) == 0 {
		return []int{-1, -1}
	}

	firstPos := binarySearch(arr, target, 0, len(arr)-1)
	if firstPos == -1 {
		return []int{-1, -1}
	}

	return []int{
		findStartPos(arr, target, firstPos),
		findEndPos(arr, target, firstPos),
	}
}

func binarySearch(arr []int, target, start, end int) int {
	for start <= end {
		mid := (start + end) / 2
		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func findStartPos(arr []int, target, firstPos int) int {
	startPos, prevStart := firstPos, firstPos

	for startPos != -1 {
		prevStart = startPos
		startPos = binarySearch(arr, target, 0, startPos-1)
	}

	return prevStart
}

func findEndPos(arr []int, target, endPos int) int {
	endPos, prevEnd := endPos, endPos

	for endPos != -1 {
		prevEnd = endPos
		endPos = binarySearch(arr, target, endPos+1, len(arr)-1)
	}

	return prevEnd
}
