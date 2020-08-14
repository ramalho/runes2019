package main

import (
	"fmt"
	"strings"

	"github.com/standupdev/runeset"
	"golang.org/x/text/unicode/runenames"
)

// CharName holds a rune and its name
type CharName struct {
	Char rune
	Name string
}

// Index maps each word to a set of runes with that word in their names
type Index map[string]runeset.Set

func (c CharName) String() string {
	return fmt.Sprintf("%U\t%c\t%s", c.Char, c.Char, c.Name)
}

func scan(start, end rune) <-chan CharName {
	ch := make(chan CharName)
	go func() {
		for r := start; r < end; r++ {
			name := runenames.Name(r)
			if len(name) > 0 && !strings.HasPrefix(name, "<") {
				ch <- CharName{r, name}
			}
		}
		close(ch)
	}()
	return ch
}

func tokenize(name string) []string {
	name = strings.Replace(name, "-", " ", -1)
	return strings.Fields(name)
}

func buildIndex(charNames <-chan CharName) (index Index) {
	index = Index{}
	for cn := range charNames {
		for _, word := range tokenize(cn.Name) {
			runes, found := index[word]
			if found {
				runes.Add(cn.Char)
			} else {
				index[word] = runeset.Make(cn.Char)
			}
		}
	}
	return index
}
