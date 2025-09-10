package thefarm

import (
    "errors"
    "fmt"
)

func DivideFood(calculator FodderCalculator, cow int) (float64, error) {
    var amount, err = calculator.FodderAmount(cow)
    if (nil != err) {
        return 0, err
    }
    var factor, err1 = calculator.FatteningFactor()
    if (nil != err1) {
        return 0, err1
    }
    return amount / float64(cow) * factor, nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, cow int) (float64, error) {
    if cow > 0 {
        return DivideFood(calculator, cow)
    }
    return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
    cow int
    message string
}
func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.cow, e.message)
}
func ValidateNumberOfCows(cow int) error {
    if cow == 0 {
        return &InvalidCowsError{
            cow: cow,
            message: "no cows don't need food",
        }
    }
    if cow < 0 {
        return &InvalidCowsError{
            cow: cow,
            message: "there are no negative cows",
        }
    }
    return nil
}