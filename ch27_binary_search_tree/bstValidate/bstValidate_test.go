package bstValidate

import (
	"challenges"
	bst "challenges/ch27_binary_search_tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	impl BstValidate
)

func TestBstValidate(t *testing.T) {
	//impl = BstValidateSolution1BruteForce{}
	//impl = BstValidateSolution2Efficient{}
	impl = BstValidateSolution2Efficient2{}

	tests := []struct {
		input         []int
		expectedValue bool
	}{
		{
			input:         []int{20, 10, 25, 5, 12, 22, 30},
			expectedValue: true,
		},
		{
			input:         []int{20, 10, 25, 5, 12, 22},
			expectedValue: true,
		},
		{
			input:         []int{20, 10, 25, 5, 12},
			expectedValue: true,
		},
		{
			input:         []int{20, 10, 25, 5},
			expectedValue: true,
		},
		{
			input:         []int{12, 7, 18, 5, 9, 16, 25},
			expectedValue: true,
		},
		{
			input:         []int{13, 6, 17, 2, 11, 10, 22},
			expectedValue: false,
		},
		{
			input:         []int{13, 6, 17, 2, 14, 10, 22},
			expectedValue: false,
		},
		{
			input:         []int{16, 8, 22, 9, 12, 19, 25},
			expectedValue: false,
		},
		{
			input:         []int{12, 7, 18, 5, 9, 14, 25, 3, 6, 10},
			expectedValue: false,
		},
		{
			input:         []int{12, 7, 18, 5, 9, 14, 25, 3, 6, 8, 11, 13, 19},
			expectedValue: false,
		},
	}

	t.Logf("Given a list of nodes in a complete binary tree, validate the same.")

	for i, tt := range tests {
		t.Logf("\ttest %d\tinput=%+v", i+1, tt.input)

		bst := &bst.BinaryNode{Data: tt.input[0]}
		bst.InsertInSpecifiedOrder(tt.input)

		actualValue := impl.validateBST(bst)
		if assert.Equal(t, tt.expectedValue, actualValue) {
			t.Logf("\t%s\texpected nodes count=%v, actual=%v", solve_problems_using_golang.Succeed, tt.expectedValue, actualValue)
		} else {
			t.Logf("\t%s\texpected nodes count=%v, actual=%v", solve_problems_using_golang.Failed, tt.expectedValue, actualValue)
		}
	}
}
