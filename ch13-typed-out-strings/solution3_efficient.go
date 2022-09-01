package ch13_typed_out_strings

type solution3_efficient struct {
}

const bk = '#' // backspace char

//{"y#fo##f", "yy#f#o##f", true},
func (sol solution3_efficient) typedOutStrings(s, t string) bool {
	// start from the end of both the strings
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		// check if char at s is #
		if (i >= 0 && s[i] == bk) || (j >= 0 && t[j] == bk) {
			if i >= 0 && s[i] == bk {
				i = skipBackspaceChars(i, s)
			}

			// repeat for t also
			if j >= 0 && t[j] == bk {
				j = skipBackspaceChars(j, t)
			}

			if i < 0 && j < 0 {
				return true
			} else if i < 0 || j < 0 {
				return false
			} else if s[i] != t[j] {
				return false
			}
			i--
			j--
		} else if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
			i, j = i-1, j-1
		} else if i >= 0 || j >= 0 {
			return false // since only one of them is present
		} else {
			return true
		}
	}

	if i < 0 && j < 0 {
		return true
	} else {
		return false
	}
}

func skipBackspaceChars(i int, in string) int {
	backCount := 2
	for backCount > 0 {
		i--
		backCount--

		if i < 0 {
			break
		}

		if in[i] == bk {
			backCount += 2
		}
	}

	return i
}
