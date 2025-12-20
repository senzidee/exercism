package bottlesong

import "strings"

var number = []string{"No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

func Recite(startBottles, takeDown int) []string {
    result := []string{}
	for i := 0; i < takeDown; i++ {
		result = append(result, sentence(startBottles)...)

        if takeDown > 1 && i < takeDown-1 {
            result = append(result, "")
        }
        startBottles--
    }
    return result
}

func sentence(startBottles int) []string {
    line1 := number[startBottles] + " green bottle" + plural(startBottles) + " hanging on the wall,"
    line2 := number[startBottles] + " green bottle" + plural(startBottles) + " hanging on the wall,"
    line3 := "And if one green bottle should accidentally fall,"
    line4 := "There'll be " + strings.ToLower(number[startBottles - 1]) + " green bottle" + plural(startBottles - 1) + " hanging on the wall."

    return []string{line1, line2, line3, line4}
}

func plural(bottle int) string {
    if bottle != 1 {
        return "s"
    }

    return ""
}
