package phonenumber

import (
    "errors"
    "fmt"
)

func Number(phoneNumber string) (string, error) {
    result := make([]rune, 0, 11)
    for _,r := range phoneNumber {
        if int(r) >= 48 && int(r) <= 57 {
            result = append(result, r)
        }
    }
    if result[0] == '1' && len(result) == 11 {
		result = result[1:]
    }
    if len(result) == 10 && result[0] >= '2' && result[0] <= '9' && result[3] >= '2' && result[3] <= '9' {
        return string(result), nil
    }

    return "", errors.New("Not a valid phone number")
}

func AreaCode(phoneNumber string) (string, error) {
	result, err := Number(phoneNumber)
    if nil != err {
        return "", err
    }

    return result[:3], nil
}

func Format(phoneNumber string) (string, error) {
	result, err := Number(phoneNumber)
    if nil != err {
        return "", err
    }

    return fmt.Sprintf("(%s) %s-%s", result[:3], result[3:6], result[6:]), err
}
