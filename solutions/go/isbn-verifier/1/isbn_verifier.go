package isbn

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	str := strings.ReplaceAll(isbn, "-", "")
	length := len(str)
	if length != 10 {
		return false
	}
	sum := 0
	for i := range str {
		if str[i] == 'X' && i == 9 {
			sum += 10 * length
		} else {
			sum += int(str[i]-'0') * length
		}
		length--
	}
	return sum%11 == 0
}
