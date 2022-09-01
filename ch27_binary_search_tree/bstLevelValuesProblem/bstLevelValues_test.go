package bstLevelValuesProblem

import (
	"challenges"
	bst "challenges/ch27_binary_search_tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	implLevelValues LevelValues
)

func TestLevelValues(t *testing.T) {
	//implLevelValues = BSTLevelValuesSolution1BruteForce{}
	implLevelValues = BSTLevelValuesSolution2Efficient{}

	tests := []struct {
		input               []int
		expectedLevelValues [][]int
	}{
		{
			input: []int{1, 2, 3, 4, 5},
			expectedLevelValues: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
		},
		{
			input: []int{20, 10, 5, 6, 7, 8, 9},
			expectedLevelValues: [][]int{
				{20},
				{10},
				{5},
				{6},
				{7},
				{8},
				{9},
			},
		},
		{
			input: []int{20, 10, 25, 5, 12, 22, 30},
			expectedLevelValues: [][]int{
				{20},
				{10, 25},
				{5, 12, 22, 30},
			},
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

		lvlVals := implLevelValues.collectLevelValues(bst)
		if assert.Equal(t, tt.expectedLevelValues, lvlVals) {
			t.Logf("\t%s\texpectedLevelValues=%v, actual=%v", solve_problems_using_golang.Succeed, tt.expectedLevelValues, lvlVals)
		} else {
			t.Logf("\t%s\texpectedLevelValues=%v, actual=%v", solve_problems_using_golang.Failed, tt.expectedLevelValues, lvlVals)
		}
	}
}
