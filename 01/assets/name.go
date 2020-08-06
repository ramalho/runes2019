package main

import (
	"fmt"

	"golang.org/x/text/unicode/runenames"
)

func main() {
	sample := "!-/Aüß♕䷕𓂀😻"
	for _, c := range sample {
		name := runenames.Name(c)
		fmt.Printf("%U\t%c\t%s\n", c, c, name)
	}
}
