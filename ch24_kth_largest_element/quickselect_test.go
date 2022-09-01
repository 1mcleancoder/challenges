package ch24_kth_largest_element

import (
	"challenges"
	"testing"
)

var (
	impl problem
)

type problem interface {
	findKthLargestElement(k int, arr []int) int
}

func TestPersonSort(t *testing.T) {
	//impl := bubbleSort1{}
	//impl := SelectionSort{}
	//impl := InsertionSort1{}
	//impl := MergeSort{}
	impl := solution1UsingQuickSort{}

	tests := []struct {
		vals     []int
		sorted   []int
		k        int
		expected int
	}{
		{
			vals:     []int{3, 7, 8, 5, 2, 1, 9, 5, 4},
			sorted:   []int{1, 2, 3, 4, 5, 5, 7, 8, 9},
			k:        3,
			expected: 7,
		},
		{
			vals:     []int{3, 7, 8, 5, 2, 1, 9, 5, 4},
			sorted:   []int{1, 2, 3, 4, 5, 5, 7, 8, 9},
			k:        1,
			expected: 9,
		},
		{
			vals:     []int{3, 7, 8, 5, 2, 1, 9, 5, 4},
			sorted:   []int{1, 2, 3, 4, 5, 5, 7, 8, 9},
			k:        5,
			expected: 5,
		},
		{
			vals:     []int{3, 7, 8, 5, 2, 1, 9, 5, 4},
			sorted:   []int{1, 2, 3, 4, 5, 5, 7, 8, 9},
			k:        4,
			expected: 5,
		},
		{
			vals:     []int{5, 3, 1, 6, 4, 2},
			sorted:   []int{1, 2, 3, 4, 5, 6},
			k:        6,
			expected: 1,
		},
		{
			vals:     []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0},
			sorted:   []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283},
			k:        1,
			expected: 283,
		},
		{
			vals:     []int{4, 2, 8, 6, 0, 5, 1, 7, 3, 9},
			sorted:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			k:        6,
			expected: 4,
		},
	}

	t.Logf("Given an array of numbers find the Kth largest number.")

	for i, tt := range tests {
		t.Logf("\ttest %d input: k=%v, arr: %+v, sorted: %+v", i, tt.k, tt.vals, tt.sorted)

		//sort.Select(Sort1UsingInbuiltSorting(tt.vals))
		//t.Logf("\nexpected:\n%+v, \nactual:\n%+v", tt.expected, tt.vals)

		actual := impl.findKthLargestElement(tt.k, tt.vals)
		if tt.expected == actual {
			t.Logf("\t%s:\t expected: %+v, actual:%+v", solve_problems_using_golang.Succeed, tt.expected, actual)
		} else {
			t.Logf("\t%s:\t expected: %+v, actual:%+v", solve_problems_using_golang.Failed, tt.expected, actual)
		}
	}
}
