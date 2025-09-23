package pangram

import "strings"

func IsPangram(input string) bool {
    for beta := 65; beta < 91; beta++ {
        if !strings.ContainsRune(strings.ToUpper(input), rune(beta)) {
            return false
        }
    }
    return true
}
