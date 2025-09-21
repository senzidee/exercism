package luhn

import (
    "strings"
	"strconv"
	"unicode"
)

func Valid(id string) bool {
	sum := 0
	runes := []rune(strings.Replace(id, " ", "", -1))
	double := false
    length := len(runes)

	for i := length - 1; i >= 0; i-- {
		r := runes[i]
		if !unicode.IsDigit(r) {
			return false
		}
		digit, _ := strconv.Atoi(string(r))
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}

	return (length > 1 && sum == 0) || (sum > 0 && sum%10 == 0)
}
