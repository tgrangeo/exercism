package anagram

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	sortedSubject := sortString(subject)
	res := []string{}

	for _, w := range candidates {
		temp := strings.ToLower(w)
		if len(subject) == len(temp) && subject != temp {
			if sortString(temp) == sortedSubject {
				res = append(res, w)
			}
		}
	}
	return res
}