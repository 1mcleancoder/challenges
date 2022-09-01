package bstBFSSearchProblem

import (
	"challenges"
	bst "challenges/ch27_binary_search_tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	bfsImpl BfsSearch
)

func TestBFSSearchHeight(t *testing.T) {
	bfsImpl = BstBFSSearchSolution1{}

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

		bst := &bst.BinaryNode{Data: tt.input[0]}
		for j, d := range tt.input {
			if j != 0 {
				bst.Insert(d)
			}
		}

		actualAllVals := bfsImpl.bfs(bst)
		if assert.Equal(t, tt.input, actualAllVals) {
			t.Logf("\t%s\tinput=%v, actual=%v", solve_problems_using_golang.Succeed, tt.input, actualAllVals)
		} else {
			t.Logf("\t%s\tinput=%v, actual=%v", solve_problems_using_golang.Failed, tt.input, actualAllVals)
		}
	}
}
