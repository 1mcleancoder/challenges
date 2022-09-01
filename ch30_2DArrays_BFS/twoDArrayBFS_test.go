package ch30_2DArrays_BFS

import (
	"challenges"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impl TwoDArrayBFS
)

func Test2DArrayDFS(t *testing.T) {
	impl = TwoDArrayBFSImpl1{}

	tests := []struct {
		input         [][]int
		startingPoint TwoD
		expectedOrder []int
	}{
		{
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
			startingPoint: TwoD{2, 2},
			expectedOrder: []int{13, 8, 14, 18, 12, 3, 9, 7, 15, 19, 17, 11, 4, 2, 10, 6, 20, 16, 5, 1},
		},
		{
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
			startingPoint: TwoD{0, 0},
			expectedOrder: []int{1, 2, 6, 3, 7, 11, 4, 8, 12, 16, 5, 9, 13, 17, 10, 14, 18, 15, 19, 20},
		},
	}

	t.Logf("Given the 2 dimentional integer array, traverse Breadth First Search: #tests=%d", len(tests))
	for i, tt := range tests {
		t.Logf("\n\ttest#%d\tinput=\n%v\n", i+1, tt.input)

		actual := impl.twoDArrayTraverseBFS(tt.input, tt.startingPoint)
		if len(tt.expectedOrder) == len(tt.expectedOrder) && assert.Equal(t, tt.expectedOrder, actual) {
			t.Logf("\t%s: bfs match: expected=%d, actual:%d", solve_problems_using_golang.Succeed, len(tt.expectedOrder), actual)
		} else {
			t.Logf("\t%s: bfs does not match: expected=%d, actual:%d", solve_problems_using_golang.Failed, len(tt.expectedOrder), actual)
		}
	}
}
