package ch3

// condition
// find len of non-repeating consecutive chars

type LongestSubstringLen2 struct {

}

// solution 1
// keep the non-repeated chars visited till now in a map so that we can find out if num is there
func (l LongestSubstringLen2) lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	idxArr := [128]int{}

	// keep the max nums count and current cons max
	var max, left, right int // sliding window
	// iterate over the string
	for i, c := range s {

		// check if num is repeated
		if idxArr[c] == 0 {
			// if not repeated then incr the current max
			idxArr[c] = i
		} else {
			if (right - left - 1) > max {
				max = right - left - 1
			}

			if idxArr[c]+1 > left {
				left = idxArr[c]
			}
			idxArr[c] = i
		}

		right++
	}

	if (right - left ) > max {
		max = right - left
	}

	return max
}

