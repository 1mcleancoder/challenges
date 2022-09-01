package ch14_longest_substring

import "math"

type solution2_efficient struct {
}

// abcdecfabecb
// abbca
func (s2 *solution2_efficient) findLongSubstringWithoutRepeatingChars(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	indexBySeenCharsMap := make(map[rune]int)

	var maxLen, currLen float64
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if idx, ok := indexBySeenCharsMap[r]; !ok {
			indexBySeenCharsMap[r] = i

			currLen++
		} else {
			maxLen = math.Max(maxLen, currLen)
			currLen = 0

			i = idx
			indexBySeenCharsMap = make(map[rune]int)
		}
	}

	return int(math.Max(maxLen, currLen))
}
