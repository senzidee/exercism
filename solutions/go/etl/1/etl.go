package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var out = make(map[string]int)
    for point,letters := range in {
        for _,letter := range letters {
        	out[strings.ToLower(letter)] = point
        }
    }

    return out
}
