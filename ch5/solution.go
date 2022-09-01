package ch5

type Problem interface {
	findLongestPalindrome(s string) string
	isPalindrome(s string) bool
}
