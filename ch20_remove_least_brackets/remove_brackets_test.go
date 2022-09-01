package ch20_remove_least_brackets

import (
	"challenges"
	"strings"
	"testing"
)

var (
	impl problem
)

func TestRemoveLeastMismatchBrackets(t *testing.T) {
	//impl = &solution1_using_stack{}
	impl = &solution2_stack_efficient{}

	tests := []struct {
		input        string
		validOutputs []string
	}{
		{"", []string{""}},
		{"a)bc(d)", []string{"abc(d)"}},
		{"(ab(c)a", []string{"ab(c)a", "(abc)a"}},
	}

	t.Logf("Given a string remove the least amount of unmatched brackets: [#tests=%d]", len(tests))

	var numFailed int
	for i, tt := range tests {
		t.Logf("\ttest: %d:\twhen input: %v", i+1, tt.input)

		actual := impl.removeLeastMismatchBrackets(tt.input)
		var found bool
		for _, ex := range tt.validOutputs {
			if actual == ex {
				found = true
				break
			}
		}

		var status string
		if !found {
			status = solve_problems_using_golang.Failed
			numFailed++
		} else {
			status = solve_problems_using_golang.Succeed
		}

		t.Logf("\t%s\t actual=%s, valid expected outputs: %+v", status, actual, tt.validOutputs)
	}

	t.Logf("Total :%s=> [#passes:%d/%d]", strings.Repeat("=", 55), len(tests)-numFailed, len(tests))
	//t.Logf("#fails:%d out of: %d %s", numFailed, len(tests), strings.Repeat("=", 30))
}
