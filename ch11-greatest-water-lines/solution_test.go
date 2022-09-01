package ch11_greatest_water_lines

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	//impl = solution1BruteForce{}
	impl = solution2{}
)

func TestFindGreatestWaterArea1(t *testing.T) {
	nums := []int{7, 1, 2, 3, 9}
	expectedArea := 7 * 4

	actualArea := impl.findGreatestWaterArea(nums)
	assert.Equalf(t, expectedArea, actualArea, "Input: %v", nums)
}

func TestFindGreatestWaterArea2(t *testing.T) {
	nums := []int{6, 9, 3, 4, 5, 8}
	expectedArea := 8 * 4

	actualArea := impl.findGreatestWaterArea(nums)
	assert.Equalf(t, expectedArea, actualArea, "Input: %v", nums)
}
