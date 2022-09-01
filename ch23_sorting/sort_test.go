package ch23_sorting

import (
	"challenges"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impl []Sort
)

type Sort interface {
	sort([]int) []int
}

func TestPersonSort(t *testing.T) {
	//impl := bubbleSort1{}
	//impl := SelectionSort{}
	//impl := InsertionSort1{}
	//impl := MergeSort{}
	impl := QuickSort{}

	tests := []struct {
		vals     []int
		expected []int
	}{
		{
			vals:     []int{3, 7, 8, 5, 2, 1, 9, 5, 4},
			expected: []int{1, 2, 3, 4, 5, 5, 7, 8, 9},
		},
		{
			vals:     []int{5, 3, 1, 6, 4, 2},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			vals:     []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0},
			expected: []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283},
		},
		{
			vals:     []int{4, 2, 8, 6, 0, 5, 1, 7, 3, 9},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	t.Logf("Given a list of Persons with fields (name, age) sort the vals by age")

	for i, tt := range tests {
		t.Logf("\ttest %d Sorting the vals: %+v", i, tt.vals)

		//sort.Sort(Sort1UsingInbuiltSorting(tt.vals))
		//t.Logf("\nexpected:\n%+v, \nactual:\n%+v", tt.expected, tt.vals)

		actual := impl.sort(tt.vals)
		if assert.Equal(t, tt.expected, tt.vals) {
			t.Logf("\t%s:\t expected: %+v, actual:%+v", solve_problems_using_golang.Succeed, tt.expected, actual)
		} else {
			t.Logf("\t%s:\t expected: %+v, actual:%+v", solve_problems_using_golang.Failed, tt.expected, actual)
		}
	}
}
