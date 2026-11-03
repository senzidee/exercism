package diamond

import (
    "errors"
    "strings"
)

func Gen(char byte) (string, error) {
    a := int(char-'A')
    if a > 26 {
        return "", errors.New("Out of range")
    }
    var b strings.Builder
    var c []string
    for i:=0;i<=a;i++ {
        b.WriteString(spaces(a - i))
        b.WriteString(string(rune('A' + i)))
        if a >= a -i +1 {
            b.WriteString(spaces((a*2+1)-(2*(a-i+1))))
            b.WriteString(string(rune('A' + i)))
        }
        b.WriteString(spaces(a - i))
        c = append(c, b.String())
        b.Reset()
    }
    for i:=a;i>0;i-- {
        c = append(c, c[i-1])
    }
    return strings.Join(c, "\n"), nil
}

func spaces(amount int) string {
    var b strings.Builder
    for i:=0;i<amount;i++ {
        b.WriteString(" ")
    }
    return b.String()
}