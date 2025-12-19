package armstrong

import (
    "math"
    "strconv"
)

func IsNumber(n int) bool {
    if n == 0 {
        return true
    }
    t := strconv.Itoa(n)
    l := len(t)
    z := 0
    for _, r := range t {
        z += int(math.Pow(float64(r - '0'), float64(l)))
    }
	return n == z
}
