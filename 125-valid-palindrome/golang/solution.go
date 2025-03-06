package golang

import "unicode"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) && left < right {
			left++
		}
		for !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) && right > left {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}
