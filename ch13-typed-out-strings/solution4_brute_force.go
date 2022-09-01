package ch13_typed_out_strings

type solution4_brute_force struct {
}

func (s4 *solution4_brute_force) typedOutStrings(s, t string) bool {
	s1 := removeBackspace(s)
	t1 := removeBackspace(t)

	if len(s1) != len(t1) {
		return false
	}

	for i := range s1 {
		if s1[i] != t1[i] {
			return false
		}
	}

	return true
}

func removeBackspace(s string) string {
	stack := make([]rune, 0, len(s))
	// iterate over s
	for _, c := range s {
		// if current char is non-# then push it
		if c != bk {
			stack = append(stack, c)
		} else if len(stack)-1 >= 0 {
			// otherwise pop a char
			stack = stack[0 : len(stack)-1]
		}
	}

	return string(stack)
}
