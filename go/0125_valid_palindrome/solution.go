package valid_palindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	r := regexp.MustCompile(`(?i)[^0-9a-z]+`)
	s = strings.ToLower(r.ReplaceAllString(s, ""))

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isPalindromeV2(s string) bool {
	const letterCaseCodesDiff = uint8(rune('a') - rune('A'))

	left, right := 0, len(s)-1

	for left < right {
		leftVal, rightVal := s[left], s[right]

		// Convert uppercase to lower
		if leftVal >= 'A' && leftVal <= 'Z' {
			leftVal += letterCaseCodesDiff
		}

		if rightVal >= 'A' && rightVal <= 'Z' {
			rightVal += letterCaseCodesDiff
		}

		// Check symbols
		if leftVal < '0' || leftVal > 'z' || (leftVal > '9' && leftVal < 'a') {
			left++
			continue
		}

		if rightVal < '0' || rightVal > 'z' || (rightVal > '9' && rightVal < 'a') {
			right--
			continue
		}

		// Compare
		if leftVal != rightVal {
			return false
		}

		left++
		right--
	}

	return true
}
