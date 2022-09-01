package ch23_sorting

type InsertionSort1 struct {
}

func (s InsertionSort1) sort(vals []int) []int {
	// 3, 0, 1, 8, 7, 2, 5, 4, 9, 6
	for i := 0; i < len(vals); i++ {
		if i >= 1 {
			if vals[i] < vals[i-1] { // last number needs to be adjusted
				for j := i; j > 0; j-- {
					if vals[j] >= vals[j-1] {
						break
					} else {
						vals[j], vals[j-1] = vals[j-1], vals[j] // swap
					}
				}
			}
		}
	}

	return vals
}
