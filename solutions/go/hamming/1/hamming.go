package hamming

import (
    "errors"
    "unicode/utf8"
)

func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
        return 0, errors.New("Can compare only if a and b have the same length")
    }
    difference := 0
    for i, _ := range a {
        if a[i] != b[i] {
            difference++
        }
    }

    return difference, nil
}
