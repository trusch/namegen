package main

import (
	"fmt"
	"time"

	"github.com/trusch/namegen/pkg/generator"
)

func main() {
	generator.Seed(time.Now().UnixNano())
	name := generator.Generate()
	fmt.Println(name)
}
