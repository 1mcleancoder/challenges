package ch13_typed_out_strings

type solution1_bf struct {
}

func (bf solution1_bf) typedOutStrings(s, t string) bool {
	s = cleanString(s)
	t = cleanString(t)

	if len(s) != len(t) {
		return false
	}

	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}

func cleanString(s string) string {
	stack := stack{
		chars: make([]rune, 0, 0),
	}
	for _, v := range s {
		if v == '#' {
			stack.pop()
		} else {
			stack.push(v)
		}
	}

	return string(stack.chars)
}

type stack struct {
	chars []rune
}

func (s *stack) push(r rune) {
	s.chars = append(s.chars, r)
}

func (s *stack) pop() rune {
	if len(s.chars) == 0 {
		return -1
	}
	lastCharIdx := len(s.chars) - 1
	c := s.chars[lastCharIdx]
	s.chars = s.chars[0:lastCharIdx]

	return c
}
