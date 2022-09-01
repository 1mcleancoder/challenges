package ch20_remove_least_brackets

type solution2_stack_efficient struct {
}

func (s2 *solution2_stack_efficient) removeLeastMismatchBrackets(in string) string {
	const oBr, cBr = '(', ')'

	// set of ints for right brackets with no left brackets
	//var cBrs = make(map[int]bool)

	var inChars = []rune(in)

	// stack of ints for left bracket indexes
	var oBrs stack
	for i, c := range inChars { // iterate over the input string
		// get the current value
		if c == oBr {
			// add index to the stack if the value is '('
			oBrs.push(i)
		} else if c == cBr {
			// if the value is ')' then check if the stack is empty
			// if the stack is empty then add the index to right bracket indexes
			// else pop the bracket
			if oBrs.isEmpty() {
				inChars[i] = 0
			} else {
				oBrs.pop() // removing it as it is valid
			}
		}
		// end of the loop
	}

	if !oBrs.isEmpty() {
		for !oBrs.isEmpty() {
			v := oBrs.pop()
			inChars[v] = 0
		}
	}

	return string(inChars)
}
