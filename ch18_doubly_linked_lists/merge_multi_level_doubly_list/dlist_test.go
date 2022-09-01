package merge_multi_level_doubly_list

import (
	"challenges"
	"testing"
)

func TestList(t *testing.T) {
	tests := []struct {
		vals     []any
		expected []any
	}{
		{
			[]any{1, 2, 2, 7, 8, 8, 10, 11, 8, 9, 2, 3, 4, 5, 5, 12, 13, 5, 6},
			[]any{1, 2, 7, 8, 10, 11, 9, 3, 4, 5, 12, 13, 6},
		},
		{
			[]any{1, 2, 2, 7, 8, 9, 2, 3, 4, 5, 5, 12, 13, 5, 6},
			[]any{1, 2, 7, 8, 9, 3, 4, 5, 12, 13, 6},
		},
	}

	t.Logf("Given a list of numbers with repeated numbers representing child nodes, flatten the list? [#tests=$%d]", len(tests))
	for i, tt := range tests {
		t.Logf("\ttest %d:\tWhen checking input=%v", i+1, tt.vals)
		list := CreateDoublyLinkedListWithChildNodes(tt.vals)
		//flattenedHead := FlattenUsingRecursion(list)
		flattenedHead := FlattenWithoutRecursion(list)
		actual := PrintWithoutChildNodes(flattenedHead)
		//actual := Print(list)

		var status string
		if len(tt.expected) == len(actual) {
			for i := range tt.expected {
				if tt.expected[i] != actual[i] {
					status = solve_problems_using_golang.Failed
					break
				} else {
					status = solve_problems_using_golang.Succeed
				}
			}
		} else {
			status = solve_problems_using_golang.Failed
		}

		if status == solve_problems_using_golang.Succeed {
			t.Logf("\t%s\tPassed with expected=actual: %v", status, actual)
		} else {
			t.Logf("\t%s\tFailed with expected:%v but got:%v", status, tt.expected, actual)
		}
	}
}
