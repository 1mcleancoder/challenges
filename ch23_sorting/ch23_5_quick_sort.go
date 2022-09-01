package ch23_sorting

type QuickSort struct {
}

func (s QuickSort) sort(arr []int) []int {
	sortRecursive(arr)
	return arr
}

func sortRecursive(arr []int) {
	p := sortPivot(arr)

	left := arr[:p]
	if len(left) > 1 {
		sortRecursive(left)
	}

	right := arr[p+1:]
	if len(right) > 1 {
		sortRecursive(right)
	}
}

func sortPivot(arr []int) int {
	if len(arr) == 1 {
		return len(arr)
	}

	p := len(arr) - 1
	i, j := 0, 0
	for j < p {
		// compare the p element to s
		if arr[j] > arr[p] {
			j++ //
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j++
		}
	} // end of loop

	arr[i], arr[p] = arr[p], arr[i]

	return i
}
