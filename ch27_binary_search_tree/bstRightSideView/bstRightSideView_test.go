package bstRightSideView

import (
	"challenges"
	bst "challenges/ch27_binary_search_tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	implRHS BstRightSideViewProblem
)

func TestLevelValues(t *testing.T) {
	//implRHS = BstRightSideViewSolution1BFS{}
	//implRHS = BstRightSideViewSolution2DFS{}
	implRHS = BstRightSideViewSolution3DFSWithArray{}

	tests := []struct {
		input       []int
		expectedRHS []int
	}{
		{
			input:       []int{20, 10, 25, 5, 12, 22, 30, 8, 6, 7},
			expectedRHS: []int{20, 25, 30, 8, 6, 7},
		},
		{
			input:       []int{20, 10, 25, 5, 12, 30, 2, 6, 7},
			expectedRHS: []int{20, 25, 30, 6, 7},
		},
	}

	for i, tt := range tests {
		t.Logf("\ttest %d\tinput=%+v", i+1, tt.input)

		bst := &bst.BinaryNode{Data: tt.input[0]}
		for j, d := range tt.input {
			if j != 0 {
				bst.Insert(d)
			}
		}

		actualRSV := implRHS.findRightSideViewValues(bst)
		if assert.Equal(t, tt.expectedRHS, actualRSV) {
			t.Logf("\t%s\texpected-right-side-view values=%v, actual=%v", solve_problems_using_golang.Succeed, tt.expectedRHS, actualRSV)
		} else {
			t.Logf("\t%s\texpected-right-side-view values=%v, actual=%v", solve_problems_using_golang.Failed, tt.expectedRHS, actualRSV)
		}
	}
}
