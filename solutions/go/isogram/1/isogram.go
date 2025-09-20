package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
    if word == "" {
        return true
    }
    m := make(map[rune]bool)
    for _, rune := range word {
        if rune == ' ' || rune == '-' || rune == '\t' || rune == '\n' {
            continue
        }
        lower := unicode.ToLower(rune)
        if m[lower] {
			return false
		}
		m[lower] = true
    }

    return true
}
