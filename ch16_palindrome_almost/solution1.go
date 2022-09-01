package ch15_palindrome_simple

import (
	"log"
	"regexp"
	"strings"
)

type solution1 struct {
}

func (sol1 *solution1) isAlmostPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	s1 := reg.ReplaceAllString(s, "")
	s2 := strings.ToLower(s1)

	//fmt.Printf("input: s=%s, s1=%s, s2=%s", s, s1, s2)

	if len(s) <= 1 {
		return true
	}

	return isAlmostPalindromeWithIndex(s2, 0, len(s2)-1, 0)
}

func isAlmostPalindromeWithIndex(s string, pL, pR, numMisses int) bool {
	// initialize left to zero and right to len-1
	// loop while left < right
	for pL < pR {
		if s[pL] != s[pR] {
			if numMisses > 0 {
				return false
			} else {
				if isAlmostPalindromeWithIndex(s, pL, pR-1, numMisses+1) || isAlmostPalindromeWithIndex(s, pL+1, pR, numMisses+1) {
					return true
				} else {
					return false
				}
			}
		}
		pL++
		pR--
	} // loop end

	return true
}
