package microblog

import (
	"strings"
)

func Truncate(phrase string) string {
	var result strings.Builder
    i := 0;
    for _, rune := range phrase {
        result.WriteRune(rune);
        i++;
        if i == 5 {
            break;
        }
    }

    return result.String()
}
