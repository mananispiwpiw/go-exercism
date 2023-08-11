package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation := '❗'
	search := '🔍'
	weather := '☀'
	var result string

	for _, value := range log {
		if value == recommendation {
			result = "recommendation"
			break
		} else if value == search {
			result = "search"
			break
		} else if value == weather {
			result = "weather"
			break
		} else {
			result = "default"
		}
	}
	return result
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var result string
	for _, value := range log {
		if value == oldRune {
			value = newRune
			result += string(value)
		} else {
			result += string(value)
		}
	}
	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit
func WithinLimit(log string, limit int) bool {
	var result bool
	if limit < utf8.RuneCountInString(log) {
		result = false
	} else if limit >= utf8.RuneCountInString(log) {
		result = true
	}
	return result
}
