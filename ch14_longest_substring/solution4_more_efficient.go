package ch14_longest_substring

type solution4_most_efficient struct {
}

func (s4 *solution4_most_efficient) findLongSubstringWithoutRepeatingChars(s string) int {
	if len(s) == 1 {
		return 1
	}

	// condition
	// find len of non-repeating consecutive chars

	// solution 1
	// keep the non-repeated chars visited till now in a map so that we can find out if num is there
	idxArr := [128]int{}
	for i := range idxArr {
		idxArr[i] = -1
	}

	// keep the max nums count and current cons max
	var max, left, right int // sliding window
	// iterate over the string
	for i, c := range s {

		// check if num is repeated
		if idxArr[c] == -1 {
			// if not repeated then incr the current max
			idxArr[c] = i
		} else {
			// calc max len
			if (right - left) > max {
				max = right - left
			}

			// reset left
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
