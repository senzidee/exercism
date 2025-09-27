package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
    if input < 1 || input > 3999 {
        return "", errors.New("Out of range")
    }
	mod := input % 10
    d := unit(mod)
    input /= 10
    mod = input % 10
    c := decimal(mod)
    input /= 10
    mod = input % 10
    b := hundreds(mod)
    input /= 10
    mod = input % 10
    a := thousands(mod)

    return a + b + c + d, nil
}
func unit(input int) string {
    switch input {
        case 1:
        	return "I"
        case 2:
        	return "II"
        case 3:
        	return "III"
        case 4:
        	return "IV"
        case 5:
        	return "V"
        case 6:
        	return "VI"
        case 7:
        	return "VII"
        case 8:
        	return "VIII"
        case 9:
        	return "IX"
        default:
        	return ""
    }
}
func decimal(input int) string {
    switch input {
        case 1:
        	return "X"
        case 2:
        	return "XX"
        case 3:
        	return "XXX"
        case 4:
        	return "XL"
        case 5:
        	return "L"
        case 6:
        	return "LX"
        case 7:
        	return "LXX"
        case 8:
        	return "LXXX"
        case 9:
        	return "XC"
        default:
        	return ""
    }
}
func hundreds(input int) string {
    switch input {
        case 1:
        	return "C"
        case 2:
        	return "CC"
        case 3:
        	return "CCC"
        case 4:
        	return "CD"
        case 5:
        	return "D"
        case 6:
        	return "DC"
        case 7:
        	return "DCC"
        case 8:
        	return "DCCC"
        case 9:
        	return "CM"
        default:
        	return ""
    }
}
func thousands(input int) string {
    switch input {
        case 1:
        	return "M"
        case 2:
        	return "MM"
        case 3:
        	return "MMM"
        default:
        	return ""
    }
}