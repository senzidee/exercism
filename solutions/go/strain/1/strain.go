package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](s []T, f func(T) bool) []T {
    r := make([]T, 0)
    for _, val := range s {
        if f(val) {
            r = append(r, val)
        }
    }

    return r
}
func Discard[T any](s []T, f func(T) bool) []T {
    r := make([]T, 0)
    for _, val := range s {
        if !f(val) {
            r = append(r, val)
        }
    }

    return r
}