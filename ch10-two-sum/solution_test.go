package ch10_two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impl solution1BruteForce
	//impl solution2UsingMaps

	nums   = []int{1, 3, 7, 9, 2}
	target = 11

	expected = []int{3, 4}
)

func TestTwoSumSolution(t *testing.T) {

	actual := impl.findTwoSum(nums, target)
	assert.NotEmpty(t, actual)
	assert.Equal(t, 2, len(actual))
	assert.EqualValues(t, expected, actual)
}

func TestTwoSumTargetSumNotPresent(t *testing.T) {
	actual := impl.findTwoSum(nums, 20)
	assert.Empty(t, actual)
}

func TestTwoSumTargetSumNoNums(t *testing.T) {
	actual := impl.findTwoSum([]int{}, 20)
	assert.Empty(t, actual)
}

func TestTwoSumTargetSumOnlyOneNum(t *testing.T) {
	actual := impl.findTwoSum([]int{20}, 20)
	assert.Empty(t, actual)
}
