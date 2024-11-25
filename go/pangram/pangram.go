package pangram

import "unicode"

func IsPangram(input string) bool {
	lettersSeen := make(map[rune]bool)
	for _, c := range input {
		if unicode.IsLetter(c) {
			letter := unicode.ToLower(c)
			lettersSeen[letter] = true
		}
	}
	return len(lettersSeen) == 26
}
