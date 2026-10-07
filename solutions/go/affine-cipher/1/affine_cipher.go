package affinecipher

import (
    "errors"
)

const M = 26

func Encode(text string, a, b int) (string, error) {
    if !coprime(mod(a, M), M) {
    	return "", errors.New("a and M must be a coprimes")
    }
    encoded := []rune{}
    n := 0
    for _, r := range text {
        c, ok := encrypt(r, a, b)
        if !ok {
        	continue
        }
        if n > 0 && n%5 == 0 {
            encoded = append(encoded, ' ')
        }
        n++
        encoded = append(encoded, c)
    }
	return string(encoded), nil
}

func Decode(text string, a, b int) (string, error) {
    aInv := modInverse(a, M)
    if aInv == -1 {
        return "", errors.New("a and M must be a coprimes")
    }
    decoded := []rune{}
    for _, r := range text {
        if c, ok := decrypt(r,aInv,a,b); ok {
            decoded = append(decoded, c)
        }
    }
    return string(decoded), nil
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }

    return a
}

func coprime(a, b int) bool {
    return gcd(a, b) == 1
}
func mod(x, m int) int {
    return ((x % m) + m) % m
}
func modInverse(a, m int) int {
    a = mod(a, m)
    for x := 1; x < m; x++ {
        if (a*x)%m == 1 {
            return x
        }
    }
    return -1
}

func encrypt(r rune, a,b int) (rune, bool) {
    switch {
    case r >= 'A' && r <= 'Z':
        return rune(mod((a * int(r - 'A') + b), M)) + 'a', true
    case r >= 'a' && r <= 'z':
        return rune(mod((a * int(r - 'a') + b), M)) + 'a', true
    case r >= '0' && r <= '9':
        return r, true
    }

     return 0, false
}

func decrypt(r rune, aInv,a,b int) (rune, bool) {
    switch {
    case r >= 'a' && r <= 'z':
        return rune(mod((aInv * (int(r - 'a') - b)), M)) + 'a', true
    case r >= '0' && r <= '9':
        return r, true
    }

     return 0, false
}