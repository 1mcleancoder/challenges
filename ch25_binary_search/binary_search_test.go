package ch25_binary_search

import (
	"challenges"
	"testing"
)

type problem interface {
	// search an element with val in the passed array; if it exists then return its index otherwise return -1
	search(arr []int, val int) int
}

var (
	impl problem
)

func TestSearch(t *testing.T) {
	impl = solution1BinarySearch{}

	tests := []struct {
		arr      []int
		val      int
		expected int
	}{
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			val:      5,
			expected: 4,
		},
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			val:      4,
			expected: 3,
		},
		{
			arr:      []int{4, 5, 6, 7, 8, 9},
			val:      9,
			expected: 5,
		},
	}

	t.Logf("Given a list of sorted integers, find the position of a particular element: #tests: %d", len(tests))

	var passed int
	for i, tt := range tests {
		t.Logf("\ttest\t%d\tarr=%+v val=%d", i, tt.arr, tt.val)
		actual := impl.search(tt.arr, tt.val)
		if actual == tt.expected {
			passed++
			t.Logf("\t%s\t\t\texpected=%v, actual=%v", solve_problems_using_golang.Succeed, tt.expected, actual)
		} else {
			t.Logf("\t%s\t\t\texpected=%v, actual=%v", solve_problems_using_golang.Failed, tt.expected, actual)
		}
	}

	t.Logf("\t#passes=%v, #failed=%v", passed, len(tests)-passed)
}
