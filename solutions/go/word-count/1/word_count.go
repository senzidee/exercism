package wordcount

import (
    "regexp"
    "strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    re := regexp.MustCompile(`[a-zA-Z]+'[a-zA-Z]+|[a-zA-Z0-9]+`)
    var fr = make(Frequency)
    for _, word := range re.FindAllString(phrase, -1) {
        key := strings.ToLower(word)
        if _, ok := fr[key]; ok {
            fr[key] += 1
        } else {
            fr[key] = 1
        }
    }
	
    return fr
}
