package ch15_palindrome_simple

import (
	"log"
	"regexp"
	"strings"
)

type solution3_bf3_compare_reversed struct {
}

func (s3 solution3_bf3_compare_reversed) isPalindrome(s string) bool {
	r, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	s = strings.ToLower(r.ReplaceAllString(s, ""))

	var sR = make([]uint8, 0, len(s))
	for i, j := len(s)-1, 0; i >= 0; i-- {
		sR = append(sR, s[i])
		j++
	}

	var mid = len(s) / 2
	for i := 0; i < mid; i++ {
		if s[i] != sR[i] {
			return false
		}
	}

	return true
}
