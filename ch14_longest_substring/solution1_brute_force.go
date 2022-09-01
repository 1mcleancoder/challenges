package ch14_longest_substring

import "math"

type solution1_bf struct {
}

// abcca
// bbbbb
func (s1 *solution1_bf) findLongSubstringWithoutRepeatingChars(s string) int {
	seenChars := make(map[rune]bool)

	var maxLen int
	for i := 0; i < len(s); i++ {
		var ssLen int
		for j := i; j < len(s); j++ {
			r := rune(s[j])
			if c, ok := seenChars[r]; !ok {
				seenChars[r] = c
				ssLen++
			} else {
				maxLen = int(math.Max(float64(maxLen), float64(ssLen)))
				ssLen = 0                       // reset
				seenChars = make(map[rune]bool) // reset seen chars
			}
		}

		maxLen = int(math.Max(float64(maxLen), float64(ssLen)))
		seenChars = make(map[rune]bool) // reset seen chars
	}

	return maxLen
}
