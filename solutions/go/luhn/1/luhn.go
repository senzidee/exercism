package luhn

import (
	"strconv"
	"unicode"
)

func Valid(id string) bool {
	sum := 0
	runes := []rune(id)
	double := false
    length := len(runes)

	for i := length - 1; i >= 0; i-- {
		r := runes[i]
		if unicode.IsSpace(r) {
			continue
		}
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

	return (length > 2 && sum == 0) || (sum > 0 && sum%10 == 0)
}
