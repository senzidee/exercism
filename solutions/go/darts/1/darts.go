package darts

func Score(x, y float64) int {
	if diagonal(x,y) <= 1.0 {
        return 10
    }
	if diagonal(x,y) <= 25.0 {
        return 5
    }
	if diagonal(x,y) <= 100.0 {
        return 1
    }
    return 0
}

func abs(x float64) float64 {
    if x < 0 {
        return x * -1
    }
    return x
}
func diagonal(x, y float64) float64 {
    return abs(x) * abs(x) + abs(y) * abs(y)
}