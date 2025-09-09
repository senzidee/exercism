package logs

import (
    "fmt"
    "unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
        switch fmt.Sprintf("%U", char) {
            case "U+2757":
            	return "recommendation"
            case "U+1F50D":
            	return "search"
            case "U+2600":
            	return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    result := ""
	for _, char := range log {
        if char == oldRune {
            result += string(newRune)
        } else {
            result += string(char)
        }
    }

    return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) > limit {
        return false
    }
    return true
}
