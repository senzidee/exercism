package summultiples

type energy struct {
    collected []int
}

func SumMultiples(limit int, divisors ...int) int {
    var e energy
    for _, divisor := range divisors {
        if divisor == 0 {
            continue
        }
        multiple := divisor
        for i := 1; multiple < limit; i++ {
            if !e.Has(multiple) {
                e.Add(multiple)
            }
            multiple += divisor
        }
    }
    return e.Sum()
}

func (e energy) Has(n int) bool {
    for _, t := range e.collected {
        if n == t {
            return true
        }
    }
    return false
}
func (e *energy) Add(n int) {
    e.collected = append(e.collected, n)
}
func (e energy) Sum() int {
    i := 0
    for _, t := range e.collected {
        i += t
    }
    return i
}