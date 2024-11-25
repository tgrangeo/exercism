package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-=~*]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`["'].*(?i)password.*["']`)
	for _, l := range lines {
		if re.MatchString(l) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	text = re.ReplaceAllString(text, "")
	return text
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(?i)\bUser\s+(\w{6,})`)
	var result []string
	for _, message := range lines {
		matches := re.FindStringSubmatch(message)
		if len(matches) > 1 {
			userName := matches[1]
			taggedMessage := fmt.Sprintf("[USR] %s %s", userName, message)
			result = append(result, taggedMessage)
		} else {
			result = append(result, message)
		}
	}

	return result
}
