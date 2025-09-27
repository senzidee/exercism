package bob

import (
    "regexp"
    "unicode"
)

func Hey(remark string) string {
    sure := regexp.MustCompile(`.\?\s*$`)
    whoa := regexp.MustCompile(`^[A-Z0-9][A-Z0-9\s!@#$%^&*(),.?:;'\-]+$`)
    calm := regexp.MustCompile(`^[A-Z0-9][A-Z\s']+\?$`)
    fine := regexp.MustCompile(`^$|^[ \s\n\r\t]+$`)
    switch {
        case calm.MatchString(remark):
        	return "Calm down, I know what I'm doing!"
        case sure.MatchString(remark):
        	return "Sure."
        case whoa.MatchString(remark) && containsUpper(remark):
        	return "Whoa, chill out!"
        case fine.MatchString(remark):
        	return "Fine. Be that way!"
        default:
        	return "Whatever."
    }
}
func containsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}