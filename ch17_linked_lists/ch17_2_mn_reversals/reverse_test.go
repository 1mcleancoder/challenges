package ch17_2_mn_reversals

import (
	"challenges"
	list "challenges/ch17_linked_lists"
	"testing"
)

var (
	//impl = solution1{}
	impl = solution2_looping_once{}
)

func TestList(t *testing.T) {
	tests := []struct {
		list         []int
		m, n         int
		expectedList string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, 5, "1,2,5,4,3,6,7"},
		{[]int{1, 2, 3, 4, 5}, 2, 4, "1,4,3,2,5"},
		{[]int{1, 2, 3, 4, 5}, 1, 5, "5,4,3,2,1"},
		{[]int{1, 2, 3, 4}, 2, 4, "1,4,3,2"},
		{[]int{1, 2, 3}, 1, 3, "3,2,1"},
		{[]int{1, 2}, 1, 2, "2,1"},
		{[]int{1, 2}, 1, 1, "1,2"},
		{[]int{1}, -1, 0, "1"},
		{[]int{}, 0, 0, ""},
	}
	n := list.Node{}

	t.Logf("Given a list of values, create its reversed list [#tests=%d]", len(tests))

	var fails, passes int
	for i, tt := range tests {
		t.Logf("\ttest %d:\tWhen reversing list: %v, m=%v, n=%v", i+1, tt.list, tt.m, tt.n)

		list := n.Create(tt.list)
		r := impl.Reverse(list, tt.m, tt.n)

		var status string
		actual := r.String()
		if actual == tt.expectedList {
			status = solve_problems_using_golang.Succeed
			passes++
		} else {
			status = solve_problems_using_golang.Failed
			fails++
		}

		t.Logf("\t%s\tShould pass when: expected:%+v == actual:%+v", status, tt.expectedList, actual)
	}

	t.Logf("\t========= #fails=%d, #passes=%d", fails, passes)
}
