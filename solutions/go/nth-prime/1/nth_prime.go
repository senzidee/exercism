package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
    if n < 1 {
        return 0, errors.New("n must be > 1")
    }
    switch n {
        case 1:
        	return 2, nil
        case 2:
        	return 3, nil
        case 3:
        	return 5, nil
        case 4:
        	return 7, nil
        case 5:
        	return 11, nil
        case 6:
        	return 13, nil
        case 7:
        	return 17, nil
        default:
        	return nth(n)
    }
}


func nth(n int) (int, error) {
    primes := []int{2, 3, 5, 7, 11, 13, 17}
    j := 19
    i := 8
    for i <= n {
        is_prime := true
        for _, prime := range primes {
            if j % prime == 0 {
                is_prime = false
            }
        }
        if is_prime {
            primes = append(primes, j)
            i++
        }
        j++
    }
    return primes[len(primes) - 1], nil
}
