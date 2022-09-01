package ch23_sorting

type InsertionSort2 struct {
}

func (s InsertionSort2) sort(vals []int) []int {
	// 3, 0, 1, 8, 7, 2, 5, 4, 9, 6
	sortedNums := make([]int, 0, len(vals))
	for i := 0; i < len(vals); i++ {
		sortedNums = append(sortedNums, vals[i])
		if len(sortedNums) > 1 {
			if sortedNums[i] < sortedNums[i-1] { // last number needs to be adjusted
				for j := len(sortedNums) - 1; j > 0; j-- {
					if sortedNums[j] >= sortedNums[j-1] {
						break
					} else {
						sortedNums[j], sortedNums[j-1] = sortedNums[j-1], sortedNums[j] // swap
					}
				}
			}
		}
	}

	return sortedNums
}
