package ch19_parenthesis_valid

type solution1 struct {
}

func (s1 *solution1) isParenthesisValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0, len(s))
	for _, v := range s {
		if isOpeningBracket(v) {
			stack = append(stack, v) // push
		} else if isClosingBracket(v) {
			if len(stack) == 0 { // there is no corresponding opening bracket, so invalid pattern
				return false
			}
			last := len(stack) - 1
			openBr := stack[last] // peek
			if !isMatching(openBr, v) {
				return false
			}

			stack = stack[0:last] // pop
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func isMatching(left, right rune) bool {
	switch left {
	case '{':
		if right == '}' {
			return true
		} else {
			return false
		}
	case '(':
		if right == ')' {
			return true
		} else {
			return false
		}
	case '[':
		if right == ']' {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func isOpeningBracket(r rune) bool {
	switch r {
	case '{', '(', '[':
		return true
	}

	return false
}

func isClosingBracket(r rune) bool {
	switch r {
	case '}', ')', ']':
		return true
	}

	return false
}
