package ch23_sorting

type SelectionSort struct {
}

func (s SelectionSort) sort(vals []int) []int {
	for i := 0; i < len(vals); i++ {

		minIdx := i
		for j := i; j < len(vals); j++ {
			if vals[minIdx] > vals[j] {
				minIdx = j
			}
		}

		vals[i], vals[minIdx] = vals[minIdx], vals[i]
	}

	return vals
}
