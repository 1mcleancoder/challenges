package ch20_remove_least_brackets

type solution1_using_stack struct {
}

func (s1 *solution1_using_stack) removeLeastMismatchBrackets(in string) string {
	const oBr, cBr = '(', ')'

	// set of ints for right brackets with no left brackets
	var cBrs = make(map[int]bool)

	// stack of ints for left bracket indexes
	var oBrs stack
	for i, c := range in { // iterate over the input string
		// get the current value
		if c == oBr {
			// add index to the stack if the value is '('
			oBrs.push(i)
		} else if c == cBr {
			// if the value is ')' then check if the stack is empty
			// if the stack is empty then add the index to right bracket indexes
			// else pop the bracket
			if oBrs.isEmpty() {
				cBrs[i] = true
			} else {
				oBrs.pop() // removing it as it is valid
			}
		}
		// end of the loop
	}

	// make a new string which would contain output
	var opt = make([]rune, 0, len(in))
	// iterate over the loop again
	for i, c := range in {
		switch c {
		case cBr:
			// if it is a closing bracket then check if the current index is among the closing brackets, if no then add into opt else do not add into opt
			if _, ok := cBrs[i]; !ok {
				opt = append(opt, c)
			}
		case oBr:
			// if it is an opening bracket then check it is among the stack elements indexes, if no then only add into output string
			if !oBrs.lookup(i) {
				opt = append(opt, c)
			}

		default:
			// if it is another char then add it into output string
			opt = append(opt, c)
		}
	}

	return string(opt)
}

type stack struct {
	vals []int
}

func (s *stack) push(in int) {
	s.vals = append(s.vals, in)
}

func (s *stack) pop() int {
	origLen := len(s.vals)
	v := s.vals[origLen-1]
	s.vals = s.vals[:origLen-1]

	return v
}

func (s *stack) isEmpty() bool {
	return len(s.vals) == 0
}
func (s *stack) lookup(val any) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}

	return false
}
