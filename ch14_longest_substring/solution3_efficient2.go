package ch14_longest_substring

import "math"

type solution3_efficient2 struct {
}

// abcdecfabecb
// abbca
// bbbbbbb
// abcabdaac
// 01234567
func (s2 *solution3_efficient2) findLongSubstringWithoutRepeatingChars(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	indexBySeenCharsMap := make(map[rune]int)

	var maxLen, currLen float64
	for left, i := 0, 0; i < len(s); i++ {
		r := rune(s[i])
		// when it is not existing or if the index is prior to the starting of the left window
		if idx, ok := indexBySeenCharsMap[r]; !ok || idx < left {
			currLen++
		} else {
			maxLen = math.Max(maxLen, currLen)
			currLen = float64(i - idx)

			left = idx
		}

		indexBySeenCharsMap[r] = i
	}

	return int(math.Max(maxLen, currLen))
}
