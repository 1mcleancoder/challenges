package ch28_priority_queue

import (
	"challenges"
	"testing"
)

var (
	impl PriorityQueue
)

func TestPQTest(t *testing.T) {
	impl = &PriorityQueueImpl1{arr: []int{}}

	tests := []struct {
		input         []int
		expectedOrder []int
	}{
		{
			input:         []int{9, 3, 6, 7, 4, 3, 8, 1, 5, 0},
			expectedOrder: []int{9, 8, 7, 6, 5, 4, 3, 3, 1, 0},
		},
		{
			input:         []int{9, 7, 4, 3, 8, 1, 5, 0},
			expectedOrder: []int{9, 8, 7, 5, 4, 3, 1, 0},
		},
	}

	t.Logf("Given the numbers find the max priority using heap: #tests=%d", len(tests))
	for i, tt := range tests {
		t.Logf("\n\ttest#%d\tadding all values in input=%v using Push()", i+1, tt.input)
		t.Logf("\t\tPush() input: #: %d, values: %v", len(tt.input), tt.input)
		for _, v := range tt.input {
			impl.Push(v)
		}

		t.Logf("\t\tverifying Len(): ")
		if impl.Len() == len(tt.input) {
			t.Logf("\t%s: len match: expected=%d, actual:%d", solve_problems_using_golang.Succeed, len(tt.input), impl.Len())
		} else {
			t.Logf("\t%s: len does not match: expected=%d, actual:%d", solve_problems_using_golang.Failed, len(tt.input), impl.Len())
		}

		t.Logf("\t\tverifying Pop(): ")
		for _, v := range tt.expectedOrder {
			aPop := impl.Pop()
			if v == aPop {
				t.Logf("\t%s: popped: expected: %v, actual: %v", solve_problems_using_golang.Succeed, v, aPop)
			} else {
				t.Logf("\t%s: popped: expected: %v, actual: %v, complete input=%v", solve_problems_using_golang.Failed, v, aPop, tt.input)
			}
		}
	}
}
