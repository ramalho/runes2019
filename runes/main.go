package main

import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/text/unicode/runenames"
)

type CharName struct {
	Char rune
	Name string
}

func (cn CharName) display() string {
	return fmt.Sprintf("%U\t%c\t%s", cn.Char, cn.Char, cn.Name)
}

func scan(start, end rune) <-chan CharName {
	output := make(chan CharName)
	go func() {
		for char := start; char < end; char ++ {
			name := runenames.Name(char)
			if len(name) > 0 && name[0] != '<' {
				output <- CharName{char, name}
			}
		}
		close(output)
	}()
	return output
}

func contains(haystack []string, needle string) bool {
	for _, s := range(haystack) {
		if s == needle {
			return true
		}
	}
	return false
}

func containsAll(haystack, needles []string) bool {
	for _, s := range(needles) {
		if ! contains(haystack, s) {
			return false
		}
	}
	return true
}

func filter(iterator <-chan CharName, query []string) <-chan CharName {
	for i, s := range query {
		query[i] = strings.ToUpper(s)
	}
	output := make(chan CharName)
	go func() {
		for cn := range iterator {
			name := strings.Fields(cn.Name)
			if containsAll(name, query) {
				output <- cn
			}
		}
		close(output)
	}()
	return output	
}

func report(words ...string) {
	charName := CharName{'\u2108', "SCRUPLE"}
	fmt.Printf("%s\n", charName.display())
	count := 1
	fmt.Printf("%d character found", count)
}

func main() {
	if len(os.Args) > 1 {
		report(os.Args[1:]...)
	} else {
		fmt.Println("Please provide one or more words to search.")
	}
}
