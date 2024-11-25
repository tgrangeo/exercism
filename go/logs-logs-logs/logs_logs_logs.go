package logs
import "strings"
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
    str := []rune(log)
	for i := 0; i < len(str);i++{
        if (str[i] == '❗'){
            return "recommendation"
        }
    	if (str[i] == '🔍'){
            return "search"
        }
    	if (str[i] == '☀'){
            return "weather"
        }
    }
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune),string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if (utf8.RuneCountInString(log) <= limit){
        return true
    }
	return false
}
