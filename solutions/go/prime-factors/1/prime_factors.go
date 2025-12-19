package prime

func Factors(n int64) []int64 {
    i := int64(2)
    r := make([]int64, 0, 0)
    for ;n > 1; {
        if (n % i == 0) {
            n /= i
            r = append(r, i)
        } else {
            i++
        }
    }

    return r
}
