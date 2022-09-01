package ch29_2DArrays_DFS

import (
	"challenges"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impl TwoDArrayDFS
)

func Test2DArrayDFS(t *testing.T) {
	impl = TwoDArrayImpl1{}

	tests := []struct {
		input         [][]int
		expectedOrder []int
	}{
		{
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
			expectedOrder: []int{1, 2, 3, 4, 5, 10, 15, 20, 19, 14, 9, 8, 13, 18, 17, 12, 7, 6, 11, 16},
		},
	}

	t.Logf("Given the 2 dimentional integer array, traverse DFS: #tests=%d", len(tests))
	for i, tt := range tests {
		t.Logf("\n\ttest#%d\tinput=\n%v\n", i+1, tt.input)

		actual := impl.twoDArrayTraverse(tt.input)
		if len(tt.expectedOrder) == len(tt.expectedOrder) && assert.Equal(t, tt.expectedOrder, actual) {
			t.Logf("\t%s: dfs match: expected=%d, actual:%d", solve_problems_using_golang.Succeed, len(tt.expectedOrder), actual)
		} else {
			t.Logf("\t%s: dfs does not match: expected=%d, actual:%d", solve_problems_using_golang.Failed, len(tt.expectedOrder), actual)
		}
	}
}
