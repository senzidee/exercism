package isbn

import (
    "regexp"
    "strconv"
    "strings"
)

func IsValidISBN(isbn string) bool {
    re := regexp.MustCompile(`^(?:ISBN(?:-10)?:? )?(?:\d{9}|(?:\d{1,5}[ -]\d{1,7}[ -]\d{1,6}[ -])?)[\dX]$`)
    if re.MatchString(isbn) {
        sum := 0
        isbn = strings.Replace(isbn, "-", "", -1)
        for i, digit := range isbn {
            if digit == 'X' {
                sum += 10
                break
            }
            num, _ := strconv.Atoi(string(digit))
            sum += num * (10 -i)
        }

        return sum % 11 == 0
    }
	return false
}
