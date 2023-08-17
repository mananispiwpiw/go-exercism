package chance

import (
	"math/rand"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	random := rand.Intn(20) + 1
	return random
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	random := rand.Float64()
	randomNumber := 0.0 + (random * 12.0)
	return randomNumber
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animal := []string{"ant",
		"beaver",
		"cat",
		"dog",
		"elephant",
		"fox",
		"giraffe",
		"hedgehog"}

	for i := len(animal) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		animal[i], animal[j] = animal[j], animal[i]
	}

	return animal
}
