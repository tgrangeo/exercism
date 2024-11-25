package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}
	for i,s := range in{
		for _,c := range s{
			res[strings.ToLower(c)] = i
		}
	}
	return res
}
