package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/unicode/runenames"
)

type CharName struct {
	Char rune
	Name string
}

func (c CharName) String() string {
	return fmt.Sprintf("%U\t%c\t%s", c.Char, c.Char, c.Name)
}

func scan(start, end rune) []CharName {
	result := []CharName{}
	for r := start; r < end; r++ {
		name := runenames.Name(r)
		if len(name) > 0 && !strings.HasPrefix(name, "<") {
			result = append(result, CharName{r, name})
		}
	}
	return result
}

func filter(sample []CharName, words []string) []CharName {
	return []CharName{}
}

func report(word string) {
	char := '\u2108'
	name := runenames.Name(char)
	fmt.Printf("%U\t%c\t%s\n", char, char, name)
	count := 1
	fmt.Printf("%d character found", count)
}

func main() {
	fmt.Println("Please provide one or more words to search.")
}
