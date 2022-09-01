package ch23_sorting

type MergeSort struct {
}

func (s MergeSort) sort(vals []int) []int {
	if len(vals) <= 1 {
		return vals
	}

	mid := len(vals) / 2
	return mergeSort(vals[:mid], vals[mid:])
}

func mergeSort(left, right []int) []int {
	if len(left) > 2 {
		mid := len(left) / 2
		left = mergeSort(left[:mid], left[mid:])
	} else if len(left) == 2 && left[0] > left[1] {
		left[0], left[1] = left[1], left[0]
	}

	if len(right) > 2 {
		mid := len(right) / 2
		right = mergeSort(right[:mid], right[mid:])
	} else if len(right) == 2 && right[0] > right[1] {
		right[0], right[1] = right[1], right[0]
	}

	return mergeArrays(left, right)
}

func mergeArrays(left, right []int) []int {
	var arr = make([]int, 0, len(left)+len(right))

	// merge the two arrays
	pL, pR := 0, 0 // pointers to left and right arrays
	for pL < len(left) && pR < len(right) {
		if left[pL] < right[pR] {
			arr = append(arr, left[pL])
			pL++
		} else {
			arr = append(arr, right[pR])
			pR++
		}
	}

	// add remaining left elements if present
	for pL < len(left) {
		arr = append(arr, left[pL])
		pL++
	}

	// add remaining right elements if present
	for pR < len(right) {
		arr = append(arr, right[pR])
		pR++
	}

	//fmt.Printf("left=%v, right=%v, merged=%v\n", left, right, arr)

	return arr
}
