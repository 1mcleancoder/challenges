package ch19_parenthesis_valid

import (
	"challenges"
	"testing"
)

var (
	impl Problem
)

func TestIsParenthesisValid(t *testing.T) {

	//impl = &solution1{}
	impl = &solution2_cleaner{}

	tests := []struct {
		input   string
		isValid bool
	}{
		{"", true},
		{"{([])}", true},
		{"{([]", false},
		{"{({])}", false},
		{"{([{}])}", true},
		{"{([{])}", false},
		{"]", false},
		{"({[]", false},
		{"[]}", false},
		{"[]{}", true},
		{"{[]()}", true},
	}

	t.Logf("Given a string with parenthesis, find whether it is valid [#tests=%d]", len(tests))

	var fails, passes int
	for i, tt := range tests {
		t.Logf("\ttest %d:\tWhen input: %v", i+1, tt.input)

		actual := impl.isParenthesisValid(tt.input)

		var status string
		if actual == tt.isValid {
			status = solve_problems_using_golang.Succeed
			passes++
		} else {
			status = solve_problems_using_golang.Failed
			fails++
		}

		t.Logf("\t%s\tShould pass when: expected:%+v == actual:%+v", status, tt.isValid, actual)
	}

	t.Logf("\t========= #fails=%d, #passes=%d", fails, passes)
}
