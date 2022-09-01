package ch3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := []string{"abcabcbb", "bbbbb", "pwwkew", "abba", " ", "aa", "au"}
	expected := []int{3, 1, 3, 2, 1, 1, 2}
	fn := LongestSubstringLen1{}
	for i := range s {
		actual := fn.lengthOfLongestSubstring(s[i])
		assert.Equalf(t, expected[i], actual, "for %q", s[i])
	}
}
