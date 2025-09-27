package clock

import "fmt"

type Clock struct {
    minutes int
}

func New(h, m int) Clock {
    c := Clock{minutes: 0}

    return c.Add(h * 60 + m)
}

func (c Clock) Add(m int) Clock {
    c.minutes += m
    c.minutes = normalize(c.minutes, 1440)
	return c
}
func normalize(m, b int) int {
    if m >= b {
        return normalize(m - b, b)
    }
    if m < 0 {
        return normalize(b + m, b)
    }
    return m
}

func (c Clock) Subtract(m int) Clock {
    c.minutes -= m
    c.minutes = normalize(c.minutes, 1440)
	return c
}

func (c Clock) String() string {
    h := normalize(c.minutes/60, 24)
	return fmt.Sprintf("%02d:%02d", h, c.minutes % 60)
}
