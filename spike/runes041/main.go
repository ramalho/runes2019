package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

type CharName struct {
	Char rune
	Name string
}

func (c CharName) String() string {
	return fmt.Sprintf("%U\t%c\t%s", c.Char, c.Char, c.Name)
}

func filterRuneNames(start, end rune, ch chan<- CharName) {
	for r := start; r < end; r++ {
		name := runenames.Name(r)
		if len(name) > 0 && !strings.HasPrefix(name, "<") {
			ch <- CharName{r, name}
		}
	}
	close(ch) // Close to end the client loop
}

func scan(start, end rune) []CharName {
	ch := make(chan CharName)
	go filterRuneNames(start, end, ch)
	result := []CharName{}
	for cn := range ch {
		result = append(result, cn)
	}
	return result
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func containsAll(haystack, needles []string) bool {
	for _, s := range needles {
		if !contains(haystack, s) {
			return false
		}
	}
	return true
}

func search(sample []CharName, words []string) []CharName {
	result := []CharName{}
	for i, s := range words {
		words[i] = strings.ToUpper(s)
	}
	for _, c := range sample {
		name := strings.Replace(c.Name, "-", " ", -1)
		parts := strings.Fields(name)
		if containsAll(parts, words) {
			result = append(result, CharName{c.Char, c.Name})
		}
	}
	return result
}

func report(words ...string) {
	result := search(scan(' ', unicode.MaxRune), words)
	for _, c := range result {
		fmt.Println(c)
	}
	fmt.Printf("%d character(s) found", len(result))
}

func main() {
	if len(os.Args) > 1 {
		report(os.Args[1:]...)
	} else {
		fmt.Println("Please provide one or more words to search.")
	}
}
