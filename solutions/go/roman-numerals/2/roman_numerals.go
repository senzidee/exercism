package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
    if input < 1 || input > 3999 {
        return "", errors.New("Out of range")
    }
    unit := [10]string{"","I","II","III","IV","V","VI","VII","VIII","IX"}
    decimal := [10]string{"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"}
    hundred := [10]string{"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"}
    thousand := [10]string{"","M","MM","MMM"}

    return thousand[(input/1000%10)] + hundred[(input/100%10)] + decimal[(input/10%10)] + unit[(input%10)], nil
}