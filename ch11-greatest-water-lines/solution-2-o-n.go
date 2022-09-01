package ch11_greatest_water_lines

import "math"

type solution2 struct {
}

// 7, 1, 2, 3, 9			min(7,9) * min(index of 9, index of 7)
// 6, 9, 3, 4, 5, 8

// 0  1  2  3  4
// 7, 1, 2, 3, 9
// lt = 0,
// rt = 5-1= 4
// w = 4, nums[0] < nums[4] => 7 < 9 => 7, a = 4 * 7 = 28, lt = 1, rt = 4

// lt = 1,
// rt = 4
// w = 3, nums[1] < nums[4] => 1 < 9 => 1, a = 3 * 1 = 3, lt = 2, rt = 4

// lt = 2,
// rt = 4
// w = 2, nums[2] < nums[4] => 2 < 9 => 2, a = 2 * 2 = 4, lt = 3, rt = 4

// lt = 3,
// rt = 4
// w = 1, nums[3] < nums[4] => 3 < 9 => 3, a = 1 * 3 = 3, lt = 4, rt = 4

// diff input:
// 0  1  2  3  4  5
// 6, 9, 3, 4, 5, 8
// lt = 0,
// rt = 6-1= 5
// w = 5, nums[0] < nums[4] => 6 < 8 => nums[0]=6, a = 5 * 6 = 30, lt = 1, rt = 5

// lt = 1,
// rt = 5
// w = 4, nums[1] < nums[4] => 9 < 8 => 8, a = 4 * 8 = 32, lt = 1, rt = 4

// lt = 1,
// rt = 4
// w = 3, nums[1] < nums[4] => 9 < 5 => 5, a = 3 * 5 = 15, lt = 1, rt = 3

// lt = 1,
// rt = 3
// w = 2, nums[1] < nums[3] => 3 < 9 => 3, a = 1 * 3 = 3, lt = 4, rt = 4

func (s solution2) findGreatestWaterArea(nums []int) int {
	var maxArea float64

	// iterate the numbers from both directions
	for lt, rt := 0, len(nums)-1; lt < rt; {
		// say lt=0 and rt=len(nums)
		// calc min(nums[lf], nums[rt]) say minHeight
		// calc diff = rt - lf say depth
		w := rt - lt // width

		var a int                 // calc area = minHeight * dapth
		if nums[lt] <= nums[rt] { // find out which index's value is minimum out of nums[lf] and nums[rt]
			a = w * nums[lt]
			lt++
		} else {
			a = w * nums[rt]
			rt--
		}

		maxArea = math.Max(maxArea, float64(a))
	}

	return int(maxArea)
}
