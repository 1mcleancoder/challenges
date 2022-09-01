package ch15_palindrome_simple

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type solution2_bf_expand_from_center struct {
}

func (sol2 *solution2_bf_expand_from_center) isPalindrome(s string) bool {
	r, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(r)
	}

	onlyAlphaNum := strings.ToLower(r.ReplaceAllString(s, ""))

	fmt.Printf("before=%s, after=%s\n", s, onlyAlphaNum)

	s = onlyAlphaNum
	// find odd or even:
	// len(s) % 2 == 0 then even else odd
	var isEven bool
	if len(s)%2 == 0 {
		isEven = true
	} else {
		isEven = false
	}

	// if odd then centerLeft = len(s) / 2
	mid := len(s) / 2

	var cL, cR int
	if isEven {
		// 0 1 2 3 => 4 => mid=2 => centerLeft = mid - 1, centerRight = mid
		cL = mid - 1
		cR = mid
	} else {
		// 0 1 2 3 4 => 5 => mid=2 => centerLeft = mid - 1, centerRight = mid + 1
		cL = mid - 1
		cR = mid + 1
	}

	// loop while centerLeft >= 0 && centerRight < len(s)
	for cL >= 0 && cR < len(s) {
		// if s[centerLeft] != s[centerRight] then return false
		if s[cL] != s[cR] {
			return false
		}
		cL--
		cR++
	} // loop end

	return true
}
