package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

const first, last = ' ', unicode.MaxRune

func main() {
	rangeNames := []string{}
	rangeCounts := make(map[string]int)
	fmt.Println("unicode.MaxRune = ", unicode.MaxRune)
	fmt.Println("runenames.UnicodeVersion = ", runenames.UnicodeVersion)
	firstNamed, lastNamed := rune(-1), unicode.MaxRune
	count := 0
	for char := first; char <= last; char++ {
		name := runenames.Name(char)
		if len(name) == 0 {
			continue
		}
		if name[0] == '<' {
			rangeCounts[name]++
			if rangeCounts[name] == 1 {
				rangeNames = append(rangeNames, name)
			}
		} else {
			count++ // names that don't start with <
			if firstNamed == -1 {
				firstNamed = char
			}
			lastNamed = char
		}
	}
	fmt.Println("Named ranges of characters (with counts):")
	for _, name := range rangeNames {
		fmt.Printf("%6d\t%s\n", rangeCounts[name], name)
	}
	fmt.Println(strings.Repeat("_", 60))
	fmt.Printf("%6d\tcharacters with unique names\n", count)
	fmt.Printf("first named character:\t%U\t%q\t%s\n", firstNamed, firstNamed, runenames.Name(firstNamed))
	fmt.Printf(" last named character:\t%U\t%q\t%s\n", lastNamed, lastNamed, runenames.Name(lastNamed))
}
