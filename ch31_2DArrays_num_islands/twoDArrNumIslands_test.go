package ch31_2DArrays_num_islands

import (
	"challenges"
	"testing"
)

var (
	impl NumIslands
)

func Test2DArrNumIslands(t *testing.T) {
	impl = TwoDArrNumIslandsImpl1{}

	tests := []struct {
		input           [][]int
		expectedIslands int
	}{
		{
			input: [][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 1},
				{0, 0, 0, 1, 1},
			},
			expectedIslands: 2,
		},
		{
			input: [][]int{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			expectedIslands: 1,
		},
		{
			input: [][]int{
				{0, 1, 0, 1, 0},
				{1, 0, 1, 0, 1},
				{0, 1, 1, 1, 0},
				{1, 0, 1, 0, 1},
			},
			expectedIslands: 7,
		},
		{
			input: [][]int{
				{0, 1, 0, 1, 0},
				{1, 0, 1, 0, 1},
				{0, 1, 0, 1, 0},
				{1, 0, 1, 0, 1},
			},
			expectedIslands: 10,
		},
	}

	t.Logf("Given the 2 dimentional integer array with 0, 1 representing water and land, find the number of islands: #tests=%d", len(tests))
	for i, tt := range tests {
		t.Logf("\n\ttest#%d\tinput=\n%v\n", i+1, tt.input)

		actualIslands := impl.findNumIslands(tt.input)
		if tt.expectedIslands == actualIslands {
			t.Logf("\t%s: num islands: expected=%d, actual:%d", challenges.Succeed, tt.expectedIslands, actualIslands)
		} else {
			t.Logf("\t%s: num islands: expected=%d, actual:%d", solve_problems_using_golang.Failed, tt.expectedIslands, actualIslands)
		}
	}
}
