package generator

import (
	"fmt"
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

type Options struct {
	Mode          Mode
	Alliteration  bool
	SubjectPrefix string
}

func Gen(opts Options) string {
	sub := chooseSubject(opts.Mode, opts.SubjectPrefix)
	adj := chooseAdjective(sub, opts.Alliteration)
	return strings.ToLower(strings.Replace(fmt.Sprintf("%s_%s", adj, sub), " ", "_", -1))
}

func chooseSubject(mode Mode, prefix string) string {
	var list []string
	switch mode {
	case NameMode:
		list = names[:]
	case AnimalMode:
		list = animals[:]
	}
	if len(prefix) > 0 {
		filtered := make([]string, 0)
		for _, sub := range list {
			sub = strings.ToLower(sub)
			if strings.HasPrefix(sub, prefix) {
				filtered = append(filtered, sub)
			}
		}
		list = filtered
	}
	return list[rand.Intn(len(list))]
}

func chooseAdjective(subject string, alliteration bool) string {
	list := adjectives[:]
	if alliteration {
		filtered := make([]string, 0)
		firstCharacter := strings.ToLower(subject[:1])
		for _, adj := range list {
			adj = strings.ToLower(adj)
			if strings.HasPrefix(adj, firstCharacter) {
				filtered = append(filtered, adj)
			}
		}
		list = filtered
	}
	return list[rand.Intn(len(list))]
}
