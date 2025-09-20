package diffsquares

import "math"

func SquareOfSum(n int) int {
	return int(math.Pow(float64(squareOfSum(n)), 2.0))
}

func squareOfSum(n int) int {
    if n == 1 {
        return n
    }
    return n + squareOfSum(n -1)
}

func SumOfSquares(n int) int {
	return sumOfSquares(n)
}

func sumOfSquares(n int) int {
    if n == 1 {
        return n
    }
    return int(math.Pow(float64(n), 2.0)) + sumOfSquares(n -1)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
