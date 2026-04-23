package eliudseggs

import "fmt"

func EggCount(displayValue int) int {
	s := fmt.Sprintf("%7b", displayValue)
    i := 0
    for _, rune := range s {
        if rune == '1' {
            i++
        }
    }

    return i
}
