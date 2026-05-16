package squareroot

import "fmt"

func SquareRoot(number int) (int, error) {
    if number < 0 {
        return 0, fmt.Errorf("cannot compute square root of negative number: %d", number)
    }
    if number == 0 {
        return 0, nil
    }
    
	guess := float64(number) / 2.0

    for range 10000 {
        next := (guess + float64(number)/guess) / 2.0
        diff := next - guess
        if diff < 0 {
            diff = -diff
        }
        if diff < 0.5 {
            result := int(next)
            if next-float64(result) >= 0.5 {
                result++
            }
            return result, nil
        }
        guess = next
    }

    result := int(guess)
    if guess-float64(result) >= 0.5 {
        result++
    }
    return result, nil
}
