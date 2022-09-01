package ch26_find_start_and_end_index

import (
	"challenges"
	"testing"
)

var (
	impl problem
)

func TestFindStartingAndEndingIndex(t *testing.T) {
	//impl = solution1LinearSearch{}
	impl = solution2BinarySearch{}

	tests := []struct {
		arr      []int // sorted input arr
		target   int
		expected []int
	}{
		{
			arr:      []int{1, 3, 3, 5, 5, 5, 8, 9},
			target:   5,
			expected: []int{3, 5},
		},
		{
			arr:      []int{1, 3, 3, 5, 5, 5, 8, 9},
			target:   3,
			expected: []int{1, 2},
		},
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 8, 9},
			target:   5,
			expected: []int{4, 4},
		},
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 8, 9},
			target:   10,
			expected: []int{-1, -1},
		},
		{
			arr:      []int{},
			target:   5,
			expected: []int{-1, -1},
		},
	}

	t.Logf("Given a sorted input array and a target value, find the staring and ending index of the target value. #tests=%d", len(tests))
	for i, tt := range tests {
		t.Logf("\ttest\t%d\tinput array=%+v, target=%v", (i + 1), tt.arr, tt.target)
		actual := impl.findStartingAndEndingIndex(tt.arr, tt.target)

		var pass bool
		if len(actual) == len(tt.expected) {
			pass = true
			for j := 0; j < len(actual); j++ {
				if actual[j] != tt.expected[j] {
					pass = false
					break
				}
			}
		}

		if pass {
			t.Logf("\t%s\t\t\texpected=%+v, actual=%+v", solve_problems_using_golang.Succeed, tt.expected, actual)
		} else {
			t.Logf("\t%s\t\t\texpected=%+v, actual=%+v", solve_problems_using_golang.Failed, tt.expected, actual)
		}
	}
}
