package ch11_greatest_water_lines

import (
	"math"
)

type solution1BruteForce struct {
}

// 7, 1, 2, 3, 9			min(7,9) * min(index of 9, index of 7)
// 6, 9, 3, 4, 5, 8
func (bf solution1BruteForce) findGreatestWaterArea(nums []int) int {
	// declare maxArea
	var maxArea float64
	for i := range nums { // iterate over the outer array from i=0 to n

		start := float64(nums[i])
		for j := i + 1; j < len(nums); j++ { // inner loop from j=i+1 to n

			minH := math.Min(start, float64(nums[j])) // calc min of nums[i] & nums[j] say minHeight
			minD := float64(j) - float64(i)           // min (i, j) say minDepth

			a := minH * minD // calc area: minHeight * minDepth say current area
			//fmt.Printf("i=%v \t j=%v, start=%v, nums[%v]=%v, minH=%v, minD=%v, area=%v, oldMaxArea=%v", i, j, start, j, nums[j], minH, minD, a, maxArea)
			maxArea = math.Max(maxArea, a) // maxArea = max(maxArea, current area)
		}
	}

	return int(maxArea)
}
