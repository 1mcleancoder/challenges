package ch4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var fn = solution2{}

func TestFindMedianSortedArrays1(t *testing.T) {
	testFindMedianSortedArrays(t, []int{1,2}, []int{3,4}, 2.5)
}

func TestFindMedianSortedArrays2(t *testing.T) {
	testFindMedianSortedArrays(t, []int{1,3}, []int{2}, 2.5)
}

func testFindMedianSortedArrays(t *testing.T, nums1, nums2 []int, expected float64) {

	actual := fn.findMedianSortedArrays(nums1, nums2)
	assert.Equalf(t, expected, actual, "for nums1=%v nums2=%v", nums1, nums2)
}