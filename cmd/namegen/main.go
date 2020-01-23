package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/trusch/namegen/pkg/generator"
)

var (
	seed         = pflag.Int64P("seed", "s", -1, "random seed (-1 is random random)")
	alliteration = pflag.BoolP("alliteration", "a", false, "generate an alliteration")
	mode         = pflag.StringP("mode", "m", "name", "generator mode ('name' or 'animal')")
)

func main() {
	pflag.Parse()
	if *seed == -1 {
		*seed = time.Now().UnixNano()
	}

	var generatorMode generator.Mode

	switch *mode {
	case "animal":
		generatorMode = generator.AnimalMode
	case "name":
		generatorMode = generator.NameMode
	default:
		logrus.Fatal("invalid mode specified")
	}

	generator.Seed(*seed)

	fmt.Println(generator.Generate(generatorMode, *alliteration))
}
