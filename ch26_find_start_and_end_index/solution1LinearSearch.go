package ch26_find_start_and_end_index

type solution1LinearSearch struct {
}

func (s1 solution1LinearSearch) findStartingAndEndingIndex(arr []int, target int) []int {
	pL, pR := 0, len(arr)-1

	start, end := -1, -1
	for pL <= pR {
		mid := (pL + pR) / 2
		if arr[mid] == target {
			start = findStart(arr, pL, mid, target)
			end = findEnd(arr, mid, pR, target)
			break
		} else if target < arr[mid] {
			pR = mid - 1
		} else {
			pL = mid + 1
		}
	}

	return []int{start, end}
}

func findStart(arr []int, start, mid, target int) int {
	reachedStart := true
	for mid >= start {
		mid--
		if arr[mid] != target {
			reachedStart = false
			break
		}
	}

	if reachedStart {
		start = mid
	} else {
		start = mid + 1
	}

	return start
}

//func findStart(arr []int, start, end, target int) int {
//	for start <= end {
//		mid := (start + end) / 2
//		if arr[mid] == target {
//			if start == mid {
//				return start
//			}
//			return findStart(arr, start, mid, target)
//		} else if target < arr[mid] {
//			end = mid - 1
//		} else {
//			start = mid + 1
//		}
//	}
//
//	return -1
//}
func findEnd(arr []int, mid, end, target int) int {
	reachedEnd := true
	for mid <= end {
		mid++
		if arr[mid] != target {
			reachedEnd = false
			break
		}
	}

	if reachedEnd {
		end = end
	} else {
		end = mid - 1
	}

	return end
}

//func findEnd(arr []int, start, end, target int) int {
//	for start <= end {
//		mid := (start + end) / 2
//		if arr[mid] == target {
//			if mid == end || arr[mid] == arr[end] {
//				return end
//			}
//			return findEnd(arr, mid, end, target)
//		} else if target < arr[mid] {
//			end = mid - 1
//		} else {
//			start = mid + 1
//		}
//	}
//
//	return -1
//}
