package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

func main() {
	rangeCounts := make(map[string]int)
	for char := '\x00'; char <= unicode.MaxRune; char++ {
		if name := runenames.Name(char); len(name) > 0 {
			rangeCounts[name]++
		}
	}
	for name, count := range rangeCounts {
		if count > 1 {
			fmt.Printf("%6d\t%s\n", count, name)
		}
	}
}
