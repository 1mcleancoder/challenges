package ch12_get_trapped_water

import "math"

type solution1BruteForce struct {
}

// height := []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
// sum = 8
func (s solution1BruteForce) getTrappedWater(heights []int) int {
	var sum int

	// iterate over the heights
	for i, v := range heights {
		// calc the max of left heights say lMax
		var lMax int
		for lt := 0; lt < i; lt++ {
			if lMax < heights[lt] {
				lMax = heights[lt]
			}
		}
		// calc the max of right heights say rMax
		var rMax int
		for rt := i + 1; rt < len(heights); rt++ {
			if rMax < heights[rt] {
				rMax = heights[rt]
			}
		}

		// calc min(lMax, rMax) say minHeight
		minH := int(math.Min(float64(lMax), float64(rMax)))
		// trappedWater = minHeight - currentValue
		tW := minH - v
		if tW > 0 {
			// add trappedWater to sum if +ve number
			sum += tW
		}
	}

	return sum
}
