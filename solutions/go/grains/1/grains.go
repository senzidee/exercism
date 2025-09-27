package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
        return 0, errors.New("Out of range")
    }
    number -= 1
    return uint64(1<<number), nil
}

func Total() uint64 {
    return uint64(1<<64 - 1)
}
