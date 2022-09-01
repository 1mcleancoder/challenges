package ch17_1_reverse_list

import (
	"challenges"
	list "challenges/ch17_linked_lists"
	"testing"
)

var (
	//impl = Solution1ReverseList{}
	impl = Solution2ReverseList{}
)

func TestList(t *testing.T) {
	tests := []struct {
		list         []int
		expectedList string
	}{
		{[]int{1, 2, 3, 4, 5}, "5,4,3,2,1"},
		{[]int{1, 2, 3, 4}, "4,3,2,1"},
		{[]int{1, 2, 3}, "3,2,1"},
		{[]int{1, 2}, "2,1"},
		{[]int{1}, "1"},
		{[]int{}, ""},
	}
	n := list.Node{}

	t.Logf("Given a list of values, create its reversed list [#tests=%d]", len(tests))

	var fails, passes int
	for i, tt := range tests {
		t.Logf("\ttest %d:\tWhen reversing list: %v", i+1, tt.list)

		list := n.Create(tt.list)
		r := impl.Reverse(list)

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
