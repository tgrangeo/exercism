package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var res strings.Builder
	for _, c := range plain {
		var newChar rune
		if unicode.IsLower(c) {
			newChar = c + int32(shiftKey)
			if newChar > 'z' {
				more := newChar - 'z'
				newChar = 'a' + more - 1
			}
		} else if unicode.IsUpper(c) {
			newChar = c + int32(shiftKey)
			if newChar > 'Z' {
				more := newChar - 'Z'
				newChar = 'A' + more - 1
			}
		} else {
			newChar = c
		}
		res.WriteRune(newChar)
	}
	return res.String()
}
