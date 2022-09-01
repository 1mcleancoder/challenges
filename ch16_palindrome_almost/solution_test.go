package ch15_palindrome_simple

import (
	"challenges"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	//impl = solution1{}
	impl = solution2{}
	//impl = solution3_bf3_compare_reversed{}
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"", true},
		{"aa", true},
		{"aa", true},
		{"ab", true},
		{"raceacar", true},
		{"abccdba", true},
		{"abccddba", false},
		{"abcdba", true},
		{"abcdefdba", false},
	}

	t.Logf("Given a string, find out whether it is a palindrome? [#tests=%d]", len(tests))

	var fails, passes int
	for i, tt := range tests {
		t.Logf("\ttest %d:\tWhen checking input=%q", i+1, tt.s)

		actual := impl.isAlmostPalindrome(tt.s)

		var status string
		if actual != tt.expected {
			status = solve_problems_using_golang.Failed
			fails++
		} else {
			status = solve_problems_using_golang.Succeed
			passes++
		}

		t.Logf("\t%s\tShould pass with isPalindrome:%t but got:%t", status, actual, tt.expected)
	}

	t.Logf("Summary: #%s: %d, #%s: %d", solve_problems_using_golang.Failed, fails, solve_problems_using_golang.Succeed, passes)
	assert.Equal(t, 0, fails)
}
