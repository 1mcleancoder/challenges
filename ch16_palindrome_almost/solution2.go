package ch15_palindrome_simple

import (
	"log"
	"regexp"
	"strings"
)

type solution2 struct {
}

func (sol1 *solution2) isAlmostPalindrome(s string) bool {
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

	var pL, pR = 0, len(s2) - 1
	for pL < pR {
		if s[pL] != s[pR] {
			return isSubPalindrome(s, pL, pR-1) || isSubPalindrome(s, pL+1, pR)
		}

		pL++
		pR--
	}
	return true
}

func isSubPalindrome(s string, pL, pR int) bool {
	// initialize left to zero and right to len-1
	// loop while left < right
	for pL < pR {
		if s[pL] != s[pR] {
			return false
		}
		pL++
		pR--
	}

	return true
}
