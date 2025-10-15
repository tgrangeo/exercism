package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	str := strings.ReplaceAll(id, " ", "")
	if len(str) <= 1 {
		return false
	}

	sum := 0
	double := false

	for i := len(str) - 1; i >= 0; i-- {
		r := rune(str[i])
		if !unicode.IsDigit(r) {
			return false
		}
		digit, _ := strconv.Atoi(string(r))
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}
