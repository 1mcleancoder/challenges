package ch15_palindrome_simple

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type solution1BfLeftRight struct {
}

func (sol1 *solution1BfLeftRight) isPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	s1 := reg.ReplaceAllString(s, "")
	s2 := strings.ToLower(s1)

	fmt.Printf("input: s=%s, s1=%s, s2=%s", s, s1, s2)

	if len(s) <= 1 {
		return true
	}

	s = s2

	// initialize left to zero and right to len-1
	pL, pR := 0, len(s)-1
	// loop while left < right
	for pL < pR {
		if s[pL] != s[pR] {
			return false
		}
		pL++
		pR--
	} // loop end

	return true
}
