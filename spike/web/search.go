package main

import (
	"strings"

	"github.com/standupdev/runeset"
)

// Search takes a runefinder.Index and a query; returns a matching set of runes.
func Search(index Index, query string) (result runeset.Set) {
	words := tokenize(strings.ToUpper(query))
	if len(words) == 0 {
		return runeset.Set{}
	}
	chars, found := index[words[0]]
	if !found {
		return runeset.Set{}
	}
	result = chars.Copy()
	for _, word := range words[1:] {
		chars, found := index[word]
		if !found {
			return runeset.Set{}
		}
		result.IntersectionUpdate(chars)
		if len(result) == 0 {
			break
		}
	}
	return result
}
