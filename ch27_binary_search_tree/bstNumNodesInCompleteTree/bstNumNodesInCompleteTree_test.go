package bstNumNodesInCompleteTree

import (
	"challenges"
	bst "challenges/ch27_binary_search_tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	findNumNodes BstNumNodesInCompleteTreeProblem
)

func TestLevelValues(t *testing.T) {
	findNumNodes = NumNodesInCompleteTreeSolution1DFS{}

	tests := []struct {
		input         []int
		expectedNodes int
	}{
		{
			input:         []int{20, 10, 25, 5, 12, 22, 30},
			expectedNodes: 7,
		},
		{
			input:         []int{20, 10, 25, 5, 12, 22},
			expectedNodes: 6,
		},
		{
			input:         []int{20, 10, 25, 5, 12},
			expectedNodes: 5,
		},
		{
			input:         []int{20, 10, 25, 5},
			expectedNodes: 4,
		},
	}

	t.Logf("Given a list of nodes in a complete binary tree, count the number of nodes")

	for i, tt := range tests {
		t.Logf("\ttest %d\tinput=%+v", i+1, tt.input)

		bst := &bst.BinaryNode{Data: tt.input[0]}
		for j, d := range tt.input {
			if j != 0 {
				bst.Insert(d)
			}
		}

		actualNodes := findNumNodes.findNumNodesInCompleteTree(bst)
		if assert.Equal(t, tt.expectedNodes, actualNodes) {
			t.Logf("\t%s\texpected nodes count=%v, actual=%v", solve_problems_using_golang.Succeed, tt.expectedNodes, actualNodes)
		} else {
			t.Logf("\t%s\texpected nodes count=%v, actual=%v", solve_problems_using_golang.Failed, tt.expectedNodes, actualNodes)
		}
	}
}
