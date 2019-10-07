package main

import (
	"fmt"

	"golang.org/x/text/unicode/runenames"
)

func main() {
	c := '氣'
	name := runenames.Name(c)
	fmt.Printf("%U\t%c\t%s\n", c, c, name)
}
