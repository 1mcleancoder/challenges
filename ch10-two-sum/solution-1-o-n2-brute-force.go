package ch10_two_sum

type solution1BruteForce struct {
}

// [1, 3, 7, 9, 2] 11
func (bf solution1BruteForce) findTwoSum(nums []int, target int) []int {
	// iterate over the numbers starting from index 0 say i=0
	for i := range nums {
		numToFind := target - nums[i]
		// create an inner loop to iterate over the number starting with index j = i + 1
		for j := i + 1; j < len(nums); j++ {
			// check if the sum of nums[i] + nums[j] == target
			if nums[j] == numToFind {
				// if yes, then return  i, j
				return []int{i, j}
			}
		}
	}
	// outside the loops return empty array
	return []int{}
}
