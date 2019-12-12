package generator

import (
	"math/rand"
	"strings"
)

// Seed seeds the math/rand generator
func Seed(seed int64) {
	rand.Seed(seed)
}

// Generate generates a new name
func Generate() string {
	i := rand.Intn(len(adjectives))
	j := rand.Intn(len(names))
	res := adjectives[i] + "_" + names[j]
	return strings.ToLower(res)
}
