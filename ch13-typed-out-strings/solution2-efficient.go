package ch13_typed_out_strings

type solution1_efficient struct {
}

// s = "bcd###",  t = "abcd###"
func (sol *solution1_efficient) typedOutStrings(s, t string) bool {
	for i, j := len(s)-1, len(t)-1; i >= 0 || j >= 0; {
		if i >= 0 && s[i] == '#' {
			cS := 1 // count starts from 1 as one # is already found

			for i--; i >= 0 && s[i] == '#'; i-- {
				cS++
			}

			// check there is a hash among the deleted chars if yes, then inc the num of counts
			c := cS - 1
			iC := i - 1
			for c > 0 {
				if iC >= 0 && s[iC] == '#' {
					cS++
				}
				c--
				iC--
			}
			i = i - cS
		}

		if j >= 0 && t[j] == '#' {
			cT := 1

			for j--; j >= 0 && t[j] == '#'; j-- {
				cT++
			}

			c := cT - 1 // since one is already checked as part of the above loop
			jC := j - 1 // counter to decrement loop
			for c > 0 {
				if jC >= 0 && t[jC] == '#' {
					cT += 2 // once for hash char and one for char to delete due to hash
				}

				c--
				jC--
			}

			j = j - cT
		}

		// when both have hashes such that all chars are escaped then strings are equal
		if i < 0 && j < 0 {
			return true
		} else if i < 0 || j < 0 {
			// when string has reached the end but the other one has chars left then decrease the corresponding index
			if i == 0 {
				if s[i] != '#' {
					return false
				}
			} else if j == 0 {
				if t[j] != '#' {
					return false
				}
			}
		} else {
			if s[i] != t[j] {
				return false
			}

			i--
			j--
		}
	}

	return true
}
