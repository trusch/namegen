package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/trusch/namegen/pkg/generator"
)

var (
	seed = flag.Int("seed", 0, "random seed to use")
)

func main() {
	flag.Parse()
	if *seed > 0 {
		generator.Seed(int64(*seed))
	} else {
		generator.Seed(time.Now().UnixNano())
	}
	name := generator.Generate()
	fmt.Println(name)
}
