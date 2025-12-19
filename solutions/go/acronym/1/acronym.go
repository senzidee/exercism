package acronym

import "strings"

func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
    s = strings.Replace(s, "_", " ", -1)
    words := strings.Fields(s)
    result := ""
    for _, word := range words {
        result += strings.ToUpper(word[:1])
    }
	return result
}
