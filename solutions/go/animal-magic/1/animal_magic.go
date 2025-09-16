package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	x := rand.Intn(20)
    if x == 0 {
        x = RollADie()
    }

    return x
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return float64(12) * rand.Float64()
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	x := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    rand.Shuffle(len(x), func(i, j int) {
    	x[i], x[j] = x[j], x[i]
    })

    return x
}
