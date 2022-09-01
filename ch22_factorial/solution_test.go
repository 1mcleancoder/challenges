package ch22_factorial

import (
	"challenges"
	"testing"
)

var (
	impl problem
)

func TestSolution(t *testing.T) {
	//impl = &solution1UsingLoops{}
	impl = &solution1UsingRecursion{}
	tests := []struct {
		input    int
		expected int
	}{
		{4, 24},
	}

	t.Logf("Given a number calculate the factorial: #tests=%d", len(tests))
	for i, tt := range tests {
		t.Logf("\ttest: %d\t calculateFactorial(%d)", i+1, tt.input)
		actual := impl.calcFactorial(tt.input)
		if tt.expected == actual {
			t.Logf("\t%s\t passes with expected: %d == actual: %d", solve_problems_using_golang.Succeed, tt.expected, actual)
		} else {
			t.Logf("\t%s\t fails with expected: %d != actual: %d", solve_problems_using_golang.Failed, tt.expected, actual)
		}
	}
}
