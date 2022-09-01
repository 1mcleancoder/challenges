package ch3

// condition
// find len of non-repeating consecutive chars

type LongestSubstringLen1 struct {

}

// solution 1
// keep the non-repeated chars visited till now in a map so that we can find out if num is there
func (l LongestSubstringLen1) lengthOfLongestSubstring(s string) int {
	idxArr := [128]int{}
	// used to store the index of the char
	for i := range idxArr {
		idxArr[i] = -1
	}

	// keep the max nums count and current cons max
	var max, left, right int // sliding window
	// iterate over the string
	for i, c := range s {

		// check if num is repeated
		if idxArr[c] == -1 {
			idxArr[c] = i // store the index of this char
		} else {
			if (right - left) > max {
				max = right - left
			}

			if idxArr[c]+1 > left {
				left = idxArr[c] + 1
			}
			idxArr[c] = i
		}

		right++
	}

	if (right - left) > max {
		max = right - left
	}

	return max
}

