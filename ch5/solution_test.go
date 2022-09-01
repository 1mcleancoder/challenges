package ch5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var fn = solution1{}

func TestSolution1(t *testing.T) {
	testSolution(t, "babad", true, "bab")
}

func TestSolution2(t *testing.T) {
	testSolution(t, "cbbd", true, "bb")
}

func testSolution(t *testing.T, in string, ex1 bool, ex2 string) {
	actual := fn.isPalindrome(in)
	assert.Equalf(t, ex1, actual, "for input: %v", in)


	//actual := fn.findLongestPalindrome(in)
	//assert.Equalf(t, expected, actual, "for input: %v", expected)
}