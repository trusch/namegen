package generator

import (
	"math/rand"
	"strings"
)

// Seed seeds the math/rand generator
func Seed(seed int64) {
	rand.Seed(seed)
}

type Mode int

const (
	NameMode Mode = iota
	AnimalMode
)

func Generate(mode Mode, alliteration bool) string {
	var (
		adjective string
		list      []string
	)
	for len(list) == 0 {
		adjective = adjectives[rand.Intn(len(adjectives))]
		switch mode {
		case NameMode:
			list = names[:]
		case AnimalMode:
			list = animals[:]
		}
		if alliteration {
			list = filterByFirstCharacter(list, adjective[0])
		}
	}
	word := list[rand.Intn(len(list))]
	result := strings.Replace(adjective+"_"+word, " ", "_", -1)
	result = strings.ToLower(result)
	return result
}

func filterByFirstCharacter(words []string, firstChar byte) (res []string) {
	for _, word := range words {
		if w := strings.ToLower(word); w[0] == firstChar {
			res = append(res, w)
		}
	}
	return res
}
