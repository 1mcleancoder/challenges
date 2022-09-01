package ch21_queue

import (
	"challenges"
	"challenges/ch21_queue/ch21_4_q_backed_by_stack"
	"testing"
)

var (
	qImpl Queue
)

func TestSimpleQ(t *testing.T) {
	//qImpl = &ch21_1_simple_q.Solution1QUsingSlices{}
	//qImpl = &ch21_2_q_backed_by_stack.SolutionQBackedByStack{}
	//qImpl = &ch21_3_q_backed_by_stack.SolutionQBackedByTwoStacks{}
	qImpl = &ch21_4_q_backed_by_stack.SolutionQBackedByTwoStacksLiterally{}

	tests := []struct {
		vals           []int
		expectedOutput []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
	}

	t.Logf("Given a queue implementation, tests its methods like Enqueue(), Dequeue(), IsEmpty(), Peek(). #tests=%d", len(tests))

	for i, tt := range tests {
		t.Logf("\t test # %d, input values are: %+v", i+1, tt.vals)

		for i, v := range tt.vals {
			beforeP := qImpl.Peek()
			qImpl.Enqueue(v)
			afterP := qImpl.Peek()
			if i == 0 {
				if beforeP != nil || afterP != v {
					t.Logf("\t %s Peek() i=%d and v=%v with beforeP = %v and afterP = %v", solve_problems_using_golang.Failed, i, v, beforeP, afterP)
				}
			} else if beforeP != afterP || beforeP != tt.vals[0] || afterP != tt.vals[0] {
				t.Logf("\t %s Peek() i=%d and v=%v with beforeP = %v and afterP = %v", solve_problems_using_golang.Failed, i, v, beforeP, afterP)
			}
		}

		var actualVals []any
		for !qImpl.IsEmpty() {
			actualVals = append(actualVals, qImpl.Dequeue())
		}

		if len(tt.expectedOutput) != len(actualVals) {
			t.Logf("\t %s with # expected values: %d [%+v] and # actual values: %d [%+v]", solve_problems_using_golang.Failed, len(tt.expectedOutput), tt.expectedOutput, len(actualVals), actualVals)
		} else {
			var fail bool
			for i, v := range tt.expectedOutput {
				if v != actualVals[i] {
					t.Logf("\t %s with expected value: [%d]=%v %+v and actual value: %v %+v", solve_problems_using_golang.Failed, i, v, tt.expectedOutput, actualVals[i], actualVals)
					fail = true
					break
				}
			}
			if !fail {
				t.Logf("\t %s with # expected values: %d %+v and # actual values: %d %+v", solve_problems_using_golang.Succeed, len(tt.expectedOutput), tt.expectedOutput, len(actualVals), actualVals)
			}
		}
	}
}
