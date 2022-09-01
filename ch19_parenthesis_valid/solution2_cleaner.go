package ch19_parenthesis_valid

type solution2_cleaner struct {
}

// "{([])}"
func (s2 *solution2_cleaner) isParenthesisValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	clBrByOpBrMap := make(map[any]any)
	clBrByOpBrMap['{'] = '}'
	clBrByOpBrMap['('] = ')'
	clBrByOpBrMap['['] = ']'

	brStack := Stack{}
	for _, v := range s {
		if _, ok := clBrByOpBrMap[v]; ok { // is it an opening bracket?
			brStack.Push(v)
		} else if brStack.IsEmpty() {
			return false // no corresponding bracket
		} else {
			openBr := brStack.Pop() // [ and v = ]
			correctClBr := clBrByOpBrMap[openBr]
			if correctClBr != v {
				return false
			}
		}
	}

	if !brStack.IsEmpty() {
		return false
	}

	return true
}

type Stack struct {
	vals []any
}

func (s *Stack) Push(v any) {
	s.vals = append(s.vals, v)
}

func (s *Stack) Pop() any {
	oldLen := len(s.vals)
	if oldLen == 0 {
		return nil
	}

	v := s.vals[oldLen-1]         // get the last element
	s.vals = s.vals[0 : oldLen-1] // shrink the slice
	return v
}

func (s *Stack) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return nil
	} else {
		return s.vals[len(s.vals)-1]
	}
}
