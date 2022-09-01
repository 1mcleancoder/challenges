package ch12_get_trapped_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	//impl = solution1BruteForce{}
	//impl = solution2{}
	impl = solution3{}
)

func TestGetTrappedWater(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
	expected := 8
	actual := impl.getTrappedWater(height)
	assert.Equalf(t, expected, actual, "height=%v", height)
}
