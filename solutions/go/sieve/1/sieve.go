package sieve

func Sieve(limit int) []int {
    primes := []int{}
    if limit < 2 {
        return primes
    }
    i := 2
    for i <= limit {
        is_prime := true
        for _, prime := range primes {
            if i % prime == 0 {
                is_prime = false
            }
        }
        if is_prime {
            primes = append(primes, i)
        }
        i++
    }
	return primes
}
