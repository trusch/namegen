package generator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	for _, mode := range []Mode{AnimalMode, NameMode} {
		for _, alliteration := range []bool{true, false} {
			for c := 'a'; c <= 'z'; c++ {
				t.Run(fmt.Sprintf("%v %v %v", mode, alliteration, string(c)), func(t *testing.T) {
					name := Gen(Options{
						Alliteration:  alliteration,
						Mode:          mode,
						SubjectPrefix: string(c),
					})
					require.NotEmpty(t, name)
				})
			}
		}
	}
}
