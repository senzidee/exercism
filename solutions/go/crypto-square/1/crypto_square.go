package cryptosquare

import (
    "math"
    "strings"
    "unicode"
)

func Encode(pt string) string {
    stripped := cleanString(pt)
    length := len(stripped)
    if length == 0 {
        return ""
    }

    t := math.Sqrt(float64(length))
    r := int(math.Round(t))
    c := r
    if t > math.Round(t) {
        c += 1
    }
    for i:= r*c - length; i > 0; i-- {
        stripped += " "
    }
    
    var builder strings.Builder
    builder.Grow(c*c)
    for i := 0; i < c; i++ {
        for j := 0; j < r; j ++ {
            builder.WriteByte(stripped[i+(j*c)])
        }
        builder.WriteByte(' ');    
    }

    result := builder.String()

	return result[:len(result)-1]
}

func cleanString(s string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsNumber(r) {
            return unicode.ToLower(r)
        }
        return -1
    }, s)
}
