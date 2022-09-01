package ch10_two_sum

type solution2UsingMaps struct {
}

//
func (s solution2UsingMaps) findTwoSum(nums []int, target int) []int {
	// create a map with key as the diff and value as the other number's index
	indexByDiffMap := make(map[int]int, len(nums))

	// iterate over the nums with i=0
	for i := range nums {
		// for each num check if the num is present in the map,
		if v, ok := indexByDiffMap[nums[i]]; ok {
			// if yes, then return its value and current index
			return []int{v, i}
		} else {
			// if no, then add the diff i.e. target - num[i] as the key into the map and value as the current index
			d := target - nums[i]
			indexByDiffMap[d] = i
		}
	}
	// outside the loop return empty array

	return []int{}
}

/** Bruteforce:
Runtime: 41 ms, faster than 20.32% of Go online submissions for Two Sum.
Memory Usage: 3.6 MB, less than 70.41% of Go online submissions for Two Sum.

T : O (n^2)
S: 	O (1)
*/

/** Using Maps:
Runtime: 3 ms, faster than 97.16% of Go online submissions for Two Sum.
Memory Usage: 4.2 MB, less than 53.82% of Go online submissions for Two Sum.

T : O (n)
S: 	O (n)

*/
