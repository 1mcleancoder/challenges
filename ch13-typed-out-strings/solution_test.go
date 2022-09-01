package ch13_typed_out_strings

import (
	"challenges"
	"testing"
)

var (
	//impl = solution1_bf{}
	//impl = solution1_efficient{}
	impl = solution3_efficient{}
	//impl = solution4_brute_force{}
)

func TestSolution(t *testing.T) {
	// table tests
	tests := []struct {
		a, b string
		e    bool
	}{
		{"bcd###", "abcd###", false},
		{"ab#z", "az#z", true},
		{"abc#d", "acc#c", false},
		{"x#y#z#", "a#", true},
		{"a###b", "#b", true},
		{"Ab#z", "ab#z", false},
		{"bxj##tw", "bxo#j##tw", true},
		{"bxj##tw", "bxj###tw", false},
		{"y#fo##f", "y#f#o##f", true},
		{"nzp#o#g", "b#nzp#o#g", true},
		{"bbbextm", "bbb#extm", false},
	}

	t.Log("Given the need to test the solution for typed out strings problem.")
	for i, tt := range tests {
		t.Logf("\tTest %d:\tWhen checking a=%q & b=%q for equality as %t", i+1, tt.a, tt.b, tt.e)
		{
			act := impl.typedOutStrings(tt.a, tt.b)
			if act == tt.e {
				t.Logf("\t%s\tShould pass with the equality as %t", solve_problems_using_golang.Succeed, tt.e)
			} else {
				//assert.Equalf(t, tt.e, act, "a=%v, b=%v expected=%v, actual=%v", tt.a, tt.b, tt.e, act)
				t.Errorf("\t%s\tShould pass with the equality as %t but got %t", solve_problems_using_golang.Failed, tt.e, act)
			}
		}
	}
}
