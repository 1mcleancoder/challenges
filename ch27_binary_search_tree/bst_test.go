package ch27_binary_search_tree

import (
	"testing"
)

var (
	impl BinaryNode
)

func TestCountHeight(t *testing.T) {

	tests := []struct {
		input               []int
		expectedHeight      int
		expectedLevelValues [][]int
	}{
		{
			input:          []int{1, 2, 3, 4, 5},
			expectedHeight: 5,
			expectedLevelValues: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
		},
		{
			input:          []int{20, 10, 5, 6, 7, 8, 9},
			expectedHeight: 7,
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
			input:          []int{20, 10, 25, 5, 12, 22, 30},
			expectedHeight: 3,
			expectedLevelValues: [][]int{
				{20},
				{10, 25},
				{5, 12, 22, 30},
			},
		}}

	for i, tt := range tests {
		t.Logf("\ttest %d\tinput=%+v, expectedHeight height=%v", i+1, tt.input, tt.expectedHeight)

		bst := &BinaryNode{Data: tt.input[0]}
		for j, d := range tt.input {
			if j != 0 {
				bst.Insert(d)
			}
		}

		//lvlVals := impl.levelValues()
		//if assert.Equal(t, tt.expectedLevelValues, lvlVals) {
		//	t.Logf("\t%s\texpectedLevelValues=%v, actual=%v", solve_problems_using_golang.Succeed, tt.expectedLevelValues, lvlVals)
		//} else {
		//	t.Logf("\t%s\texpectedLevelValues=%v, actual=%v", solve_problems_using_golang.Failed, tt.expectedLevelValues, lvlVals)
		//}
	}
}
