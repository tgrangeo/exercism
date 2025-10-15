package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(str string) Frequency {
	str = strings.ToLower(str)
	re := regexp.MustCompile(`\b[a-zA-Z0-9]+('[a-zA-Z0-9]+)?\b`)
	words := re.FindAllString(str, -1)
	wordCountMap := make(map[string]int)
	for _, word := range words {
		wordCountMap[word]++
    }
	return wordCountMap
}