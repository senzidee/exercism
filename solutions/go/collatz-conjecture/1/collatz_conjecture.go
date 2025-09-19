package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
		return 0, errors.New("the number must be positive")
	}
    iterated := 0
	for n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = n * 3 + 1
        }
        iterated++
    }

    return iterated, nil
}
