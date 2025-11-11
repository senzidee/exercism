package lsproduct

import (
    "errors"
    "strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if len(digits) < span {
        return 0, errors.New("span must be smaller than string length")
    }
    if span < 0 {
        return 0, errors.New("span must not be negative")
    }
    runes := []rune(digits)
    big := int64(0)
    for i := 0; i <= len(runes) - span; i++ {
        cur := int64(1)
        for _, c := range runes[i:i+span] {
            j, err := strconv.Atoi(string(c))
            if nil != err {
                return 0, err
            }
            cur *= int64(j)
        }
        if cur > big {
            big = cur
        }
    }

    return big, nil
}
