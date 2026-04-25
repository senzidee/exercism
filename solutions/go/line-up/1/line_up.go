package lineup

import "fmt"

func Format(name string, number int) string {
    suffix := "th"
    switch number % 10 {
        case 1:
        suffix = "st"
        case 2:
        suffix = "nd"
        case 3:
        suffix = "rd"
    }
    if number > 100 || number > 10 && number < 14 {
        switch number % 100 {
            case 11:
            suffix = "th"
            case 12:
            suffix = "th"
            case 13:
            suffix = "th"
        }
    }
	return fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, suffix)
}
