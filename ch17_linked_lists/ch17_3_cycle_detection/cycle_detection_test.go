package ch17_3_cycle_detection

import (
	"challenges"
	list "challenges/ch17_linked_lists"
	"testing"
)

var (
	impl Problem
)

func TestList(t *testing.T) {
	//impl = solution1_bf{}
	impl = &solution2EfficientFloydsTortoiseHareAlgo{}

	tests := []struct {
		list             []int
		expectedCycleVal int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 3}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 2}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 6}, 6},
	}
	n := list.Node{}

	t.Logf("Given a list of values, create its reversed list [#tests=%d]", len(tests))

	var fails, passes int
	for i, tt := range tests {
		t.Logf("\ttest %d:\tWhen detecting cycle in list: %v", i+1, tt.list)

		list := n.CreateListWithCycleWhenSimilarValues(tt.list)
		cycleNode := impl.detectCycle(list)

		var status string
		if cycleNode == nil && tt.expectedCycleVal == 0 {
			status = solve_problems_using_golang.Succeed
			passes++
			t.Logf("\t%s\tShould pass when: expected:%+v == actual:%+v", status, tt.expectedCycleVal, cycleNode)
		} else if cycleNode != nil && cycleNode.Val == tt.expectedCycleVal {
			status = solve_problems_using_golang.Succeed
			passes++

			t.Logf("\t%s\tShould pass when: expected:%+v == actual:%+v", status, tt.expectedCycleVal, cycleNode.Val)
		} else {
			status = solve_problems_using_golang.Failed
			fails++

			if cycleNode == nil {
				t.Logf("\t%s\tShould pass when: expected:%+v == actual:%+v", status, tt.expectedCycleVal, nil)
			} else {
				t.Logf("\t%s\tShould pass when: expected:%+v == actual:%+v", status, tt.expectedCycleVal, cycleNode.Val)
			}
		}
	}

	t.Logf("\t========= #fails=%d, #passes=%d", fails, passes)
}
