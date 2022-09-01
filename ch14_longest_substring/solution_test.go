package ch14_longest_substring

import (
	"challenges"
	"testing"
)

var (
	//impl = solution1_bf{}
	//impl = solution2_efficient{}
	//impl = solution3_efficient2{}
	impl = solution4_most_efficient{}
)

func TestSolutions(t *testing.T) {
	tests := []struct {
		input                       string
		expectedLongestSubstringLen int
	}{
		{"abcbdaac", 4},
		{"abccabb", 3},
		{"bbbbbbb", 1},
		{"", 0},
		{"abbca", 3},
		{"abcbda", 4},
		{"abcadeb", 5},
	}

	t.Logf("Given the input strings, find the longest substrings lengths [#tests=%d]", len(tests))

	var fails, passes int
	for i, tt := range tests {
		t.Logf("\tTest %d:\twhen checking input=%q", i+1, tt.input)

		actual := impl.findLongSubstringWithoutRepeatingChars(tt.input)
		if actual != tt.expectedLongestSubstringLen {
			t.Logf("\t%s\tShould pass with the longest string as %d but got %d", solve_problems_using_golang.Failed, tt.expectedLongestSubstringLen, actual)
			fails++
		} else {
			t.Logf("\t%s\tShould pass with the longest string as %d but got %d", solve_problems_using_golang.Succeed, tt.expectedLongestSubstringLen, actual)
			passes++
		}
	}

	t.Logf("#%s: %d, #%s: %d", solve_problems_using_golang.Failed, fails, solve_problems_using_golang.Succeed, passes)
}
