package scrabble

import "strings"

var letterGroups = map[int]string{
	1:  "AEIOULNRST",
	2:  "DG",
	3:  "BCMP",
	4:  "FHVWY",
	5:  "K",
	8:  "JX",
	10: "QZ",
}

func Score(word string) int {
	totalScore := 0
	for _, letter := range strings.ToUpper(word) {
		for score, letters := range letterGroups {
			if strings.ContainsRune(letters, letter) {
				totalScore += score
				break
			}
		}
	}
	return totalScore
}