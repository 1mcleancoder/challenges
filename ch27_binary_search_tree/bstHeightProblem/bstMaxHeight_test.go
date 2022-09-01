package bstHeightProblem

import (
	"challenges"
	"challenges/ch27_binary_search_tree"
	"testing"
)

var (
	maxHeightImpl HeightProblem
)

func TestFindCountHeight(t *testing.T) {
	maxHeightImpl = MaxHeightSolution1BruteForce{}

	tests := []struct {
		input               []int
		expectedHeight      int
		expectedLevelValues [][]int
	}{
		{
			input:          []int{1, 2, 3, 4, 5},
			expectedHeight: 5,
		},
		{
			input:          []int{20, 10, 5, 6, 7, 8, 9},
			expectedHeight: 7,
		},
		{
			input:          []int{20, 10, 25, 5, 12, 22, 30},
			expectedHeight: 3,
		},
	}

	for i, tt := range tests {
		t.Logf("\ttest %d\tinput=%+v, expectedHeight height=%v", i+1, tt.input, tt.expectedHeight)

		bst := &ch27_binary_search_tree.BinaryNode{Data: tt.input[0]}
		for j, d := range tt.input {
			if j != 0 {
				bst.Insert(d)
			}
		}

		actual := maxHeightImpl.findMaxHeight(bst)
		if tt.expectedHeight == actual {
			t.Logf("\t%s\texpectedHeight=%v, actual=%v", solve_problems_using_golang.Succeed, tt.expectedHeight, actual)
		} else {
			t.Logf("\t%s\texpectedHeight=%v, actual=%v", solve_problems_using_golang.Failed, tt.expectedHeight, actual)
		}
	}
}
