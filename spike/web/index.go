package main

import (
	"fmt"
	"strings"

	"github.com/standupdev/runeset"
	"golang.org/x/text/unicode/runenames"
)

// RuneName holds a rune and its name
type RuneName struct {
	Code rune
	Name string
}

// Index maps each word to a set of runes with that word in their names
type Index map[string]runeset.Set

func (c RuneName) String() string {
	return fmt.Sprintf("%U\t%c\t%s", c.Code, c.Code, c.Name)
}

func scan(start, end rune) <-chan RuneName {
	ch := make(chan RuneName)
	go func() {
		for r := start; r < end; r++ {
			name := runenames.Name(r)
			if len(name) > 0 && !strings.HasPrefix(name, "<") {
				ch <- RuneName{r, name}
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

func buildIndex(RuneNames <-chan RuneName) (index Index) {
	index = Index{}
	for cn := range RuneNames {
		for _, word := range tokenize(cn.Name) {
			runes, found := index[word]
			if found {
				runes.Add(cn.Code)
			} else {
				index[word] = runeset.Make(cn.Code)
			}
		}
	}
	return index
}
