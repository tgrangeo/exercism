package isogram

import "strings"

func isAlphaNumeric(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func IsIsogram(word string) bool {
	used := ""
	for _, c := range strings.ToUpper(word) {
		if strings.Contains(used, string(c)) {
			return false
		}
		if isAlphaNumeric(c) {
			used += string(c)
		}
	}
	return true
}
